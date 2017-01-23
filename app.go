package pub

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/tylerb/graceful"
)

// Start ...
func Start(credential *Credential, config *Config) ExitCode {
	logrus.Info("App will start")

	// FIXME 各種ハンドラー生成処理やロガー初期化処理等をnewCtxに逃がす！
	ctx, err := newCtx(credential, config)
	if err != nil {
		logrus.Errorf("[Start][call newCtx()] %#v\n", err)
		return ExitCodeCtxError
	}
	fmt.Println(ctx)

	// FIXME ↓以降を適宜、↑のnewCtxに逃がす！

	awsHandler, err := newAwsHandler(config.aws, credential)
	if err != nil {
		logrus.Errorf("AWS setting error %#v", err)
		return ExitCodeAwsSettingError
	}
	logrus.Info("AWS connect setting done")

	lineHandler, err := newLineHandler(config.line, credential)
	if err != nil {
		logrus.Errorf("LINE setting error: %#v", err)
		return ExitCodeLineSettingError
	}
	logrus.Info("LINE connect setting done")

	mux := http.NewServeMux()

	handler := &webHandler{
		config:      config,
		awsHandler:  awsHandler,
		lineHandler: lineHandler,
	}

	mux.HandleFunc(config.line.webhookURL, handler.HandlerFunc)

	logrus.Infof("Server will start at Port[%s]", config.server.port)
	graceful.Run(config.server.port, 1*time.Second, mux)
	logrus.Infof("Server stop at Port[%s]", config.server.port)

	return ExitCodeOK
}
