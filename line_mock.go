package pub

import "github.com/line/line-bot-sdk-go/linebot"

type lineHandlerMock struct {
}

func (h *lineHandlerMock) getClient() *linebot.Client {
	return nil
}
