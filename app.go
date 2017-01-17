package pub

import (
	"net/http"
	"time"

	"github.com/tylerb/graceful"
)

// App ...
type App struct {
	logger      *appLogger
	config      *Config
	awsHandler  awsHandlerIF
	lineHandler lineHandlerIF
}

// NewApp ...
func NewApp(credential *Credential, config *Config) (*App, ExitCode) {
	logger, err := newAppLogger(config.logger)
	if err != nil {
		return nil, ExitCodeLogSetupError
	}

	awsHandler, err := newAwsHandler(config.aws, credential, logger)
	if err != nil {
		logger.entry.Errorf("AWS setting error %#v", err)
		return nil, ExitCodeAwsSettingError
	}
	logger.entry.Info("AWS connect setting done")

	lineHandler, err := newLineHandler(config.line, credential, logger)
	if err != nil {
		logger.entry.Error("LINE setting error: %#v", err)
		return nil, ExitCodeLineSettingError
	}
	logger.entry.Info("LINE connect setting done")

	app := &App{
		logger:      logger,
		config:      config,
		awsHandler:  awsHandler,
		lineHandler: lineHandler,
	}
	return app, ExitCodeOK
}

// Start ...
func (a *App) Start() ExitCode {
	a.logger.entry.Info("App will start")

	mux := http.NewServeMux()

	handler := &webHandler{
		logger:      a.logger,
		config:      a.config,
		awsHandler:  a.awsHandler,
		lineHandler: a.lineHandler,
	}

	mux.HandleFunc(a.config.line.webhookURL, handler.HandlerFunc)

	a.logger.entry.Infof("Server will start at Port[%s]", a.config.server.port)
	graceful.Run(a.config.server.port, 1*time.Second, mux)
	a.logger.entry.Infof("Server stop at Port[%s]", a.config.server.port)

	return ExitCodeOK
}

// Close ...
func (a *App) Close() error {
	return a.logger.Close()
}
