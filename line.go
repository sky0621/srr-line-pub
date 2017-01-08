package pub

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

type lineHandlerIF interface {
	getClient() *linebot.Client
	parseRequest(r *http.Request) ([]linebot.Event, error)
	validateSignature(signature string, body []byte) bool
}

type lineHandler struct {
	client *linebot.Client
	token  string
	secret string
	logger *appLogger
}

func newLineHandler(cfg *lineConfig, arg *Arg, logger *appLogger) (lineHandlerIF, error) {
	if cfg.environment == constEnvLocal {
		return &lineHandlerMock{}, nil
	}

	cli, err := linebot.New(arg.lineChannelSecret, arg.lineAccessToken)
	if err != nil {
		return nil, err
	}
	return &lineHandler{
		client: cli,
		token:  arg.lineAccessToken,
		secret: arg.lineChannelSecret,
		logger: logger,
	}, nil
}

func (h *lineHandler) getClient() *linebot.Client {
	return h.client
}

func (h *lineHandler) parseRequest(r *http.Request) ([]linebot.Event, error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if !h.validateSignature(r.Header.Get("X-LINE-Signature"), body) {
		h.logger.entry.Warn(r.Header.Get("X-LINE-Signature"))
	}

	request := &struct {
		Events []linebot.Event `json:"events"`
	}{}
	if err = json.Unmarshal(body, request); err != nil {
		return nil, err
	}
	return request.Events, nil
}

func (h *lineHandler) validateSignature(signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	h.logger.entry.Debug("##### validateSignature #####")
	h.logger.entry.Debug(decoded)
	hash := hmac.New(sha256.New, []byte(h.secret))
	hash.Write(body)
	h.logger.entry.Debug(hash)
	return hmac.Equal(decoded, hash.Sum(nil))
}
