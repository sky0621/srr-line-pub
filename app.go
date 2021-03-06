package pub

import (
	"net/http"
	"time"

	"github.com/sky0621/go-lib/log"
	"github.com/sky0621/srr-line-pub/global"
	"github.com/sky0621/srr-line-pub/static"
	"github.com/tylerb/graceful"
)

// StartApp ...
func StartApp(ctx *Ctx) static.ExitCode {
	global.L.Log(log.I, "App will start")

	mux := http.NewServeMux()

	handler := &webHandler{
		awsHandler:  ctx.awsHandler,
		lineHandler: ctx.lineHandler,
	}

	mux.HandleFunc(ctx.config.line.webhookURL, handler.HandlerFunc)

	global.L.Logf(log.I, "Server will start at Port[%s]", ctx.config.server.port)
	graceful.Run(ctx.config.server.port, 1*time.Second, mux)
	global.L.Logf(log.I, "Server stop at Port[%s]", ctx.config.server.port)

	return static.ExitCodeOK
}
