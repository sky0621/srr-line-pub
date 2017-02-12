package pub

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/sky0621/srr-line-pub/static"
	"github.com/tylerb/graceful"
)

// StartApp ...
func StartApp(ctx *Ctx) static.ExitCode {
	logrus.Info("App will start")

	mux := http.NewServeMux()

	handler := &webHandler{
		awsHandler:  ctx.awsHandler,
		lineHandler: ctx.lineHandler,
	}

	mux.HandleFunc(ctx.config.line.webhookURL, handler.HandlerFunc)

	logrus.Infof("Server will start at Port[%s]", ctx.config.server.port)
	graceful.Run(ctx.config.server.port, 1*time.Second, mux)
	logrus.Infof("Server stop at Port[%s]", ctx.config.server.port)

	return static.ExitCodeOK
}
