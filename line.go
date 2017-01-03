package pub

import "github.com/line/line-bot-sdk-go/linebot"

type lineHandlerIF interface {
	getClient() *linebot.Client
}

type lineHandler struct {
	client *linebot.Client
}

func newLineHandler(cfg *lineConfig, arg *Arg) (lineHandlerIF, error) {
	if cfg.environment == constEnvLocal {
		return &lineHandlerMock{}, nil
	}

	cli, err := linebot.New(arg.lineChannelSecret, arg.lineAccessToken)
	if err != nil {
		return nil, err
	}
	return &lineHandler{client: cli}, nil
}

func (h *lineHandler) getClient() *linebot.Client {
	return h.client
}
