package pub

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/uber-go/zap"
)

type webHandler struct {
	ctx *ctx
}

func (h *webHandler) HandlerFunc(c echo.Context) error {
	events, err := h.ctx.lineCli.ParseRequest(c.Request())
	if err != nil {
		h.ctx.logger.entry.Error("error: %#v", zap.Error(err))
		return err
	}

	for _, event := range events {
		h.ctx.logger.entry.Info(fmt.Sprintf("event: %#v", event))
		// &sqs.
	}

	// FIXME events を処理する。
	// FIXME aws-sdk-go/sqs にてキューに投入

	return nil
}
