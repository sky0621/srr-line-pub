package pub

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/tylerb/graceful"
)

// App ...
type App struct {
	config      *Config
	awsHandler  awsHandlerIF
	lineHandler lineHandlerIF
}

// NewApp ...
func NewApp(credential *Credential, config *Config) (*App, ExitCode) {
	awsHandler, err := newAwsHandler(config.aws, credential)
	if err != nil {
		logrus.Errorf("AWS setting error %#v", err)
		return nil, ExitCodeAwsSettingError
	}
	logrus.Info("AWS connect setting done")

	lineHandler, err := newLineHandler(config.line, credential)
	if err != nil {
		logrus.Errorf("LINE setting error: %#v", err)
		return nil, ExitCodeLineSettingError
	}
	logrus.Info("LINE connect setting done")

	app := &App{
		config:      config,
		awsHandler:  awsHandler,
		lineHandler: lineHandler,
	}
	return app, ExitCodeOK
}

// Start ...
func (a *App) Start() ExitCode {
	logrus.Info("App will start")

	mux := http.NewServeMux()

	handler := &webHandler{
		config:      a.config,
		awsHandler:  a.awsHandler,
		lineHandler: a.lineHandler,
	}

	mux.HandleFunc(a.config.line.webhookURL, handler.HandlerFunc)

	logrus.Infof("Server will start at Port[%s]", a.config.server.port)
	graceful.Run(a.config.server.port, 1*time.Second, mux)
	logrus.Infof("Server stop at Port[%s]", a.config.server.port)

	return ExitCodeOK
}

// Close ...
func (a *App) Close() error {
	return nil
}
