package pub

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

type lineHandlerMock struct {
}

func (h *lineHandlerMock) getClient() *linebot.Client {
	return nil
}

func (h *lineHandlerMock) parseRequest(r *http.Request) ([]linebot.Event, error) {
	return nil, nil
}

func (h *lineHandlerMock) validateSignature(signature string, body []byte) bool {
	return true
}
