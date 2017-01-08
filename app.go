package pub

import (
	"net/http"
	"time"

	"github.com/tylerb/graceful"
)

// App ...
type App struct {
	logger      *appLogger
	config      *config
	awsHandler  awsHandlerIF
	lineHandler lineHandlerIF
}

// NewApp ...
func NewApp(arg *Arg) (*App, int) {
	err := readConfig(arg.configFilePath)
	if err != nil {
		panic(err)
	}

	logger, err := newAppLogger(newLoggerConfig())
	if err != nil {
		return nil, ExitCodeLogSetupError
	}

	awsHandler, err := newAwsHandler(newAwsConfig())
	if err != nil {
		logger.entry.Errorf("AWS setting error %#v", err)
		return nil, ExitCodeAwsSettingError
	}
	logger.entry.Info("AWS connect setting done")

	lineHandler, err := newLineHandler(newLineConfig(), arg)
	if err != nil {
		logger.entry.Error("LINE setting error: %#v", err)
		return nil, ExitCodeLineSettingError
	}
	logger.entry.Info("LINE connect setting done")

	app := &App{
		logger:      logger,
		config:      newConfig(),
		awsHandler:  awsHandler,
		lineHandler: lineHandler,
	}
	return app, ExitCodeOK
}

// Start ...
func (a *App) Start() int {
	a.logger.entry.Info("App will start")

	// e := webSetup(a.config.logger)
	mux := http.NewServeMux()

	handler := &webHandler{
		logger:      a.logger,
		config:      a.config,
		awsHandler:  a.awsHandler,
		lineHandler: a.lineHandler,
	}

	// e.POST(a.config.line.webhookURL, handler.HandlerFunc)
	mux.HandleFunc(a.config.line.webhookURL, handler.HandlerFunc)

	a.logger.entry.Infof("Server will start at Port[%s]", a.config.server.port)
	// e.Logger.Infof("Echo Server will start at Port[%s]", a.config.server.port)
	// err := e.Start(a.config.server.port)
	graceful.Run(a.config.server.port, 1*time.Second, mux)
	// if err != nil {
	// 	a.logger.entry.Error("error: %#v", err)
	// 	return ExitCodeServerStartError
	// }

	return ExitCodeOK
}

// Close ...
func (a *App) Close() error {
	return a.logger.Close()
}
