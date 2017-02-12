package pub

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sky0621/srr-line-pub/static"
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
}

func newLineHandler(cfg *lineConfig, credential *Credential) (lineHandlerIF, error) {
	if cfg.environment == static.ConstEnvLocal {
		return &lineHandlerMock{}, nil
	}

	cli, err := linebot.New(credential.LineChannelSecret, credential.LineAccessToken)
	if err != nil {
		return nil, err
	}
	return &lineHandler{
		client: cli,
		token:  credential.LineAccessToken,
		secret: credential.LineChannelSecret,
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
		logrus.Warn(r.Header.Get("X-LINE-Signature"))
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
	logrus.Debug("##### validateSignature #####")
	logrus.Debug(decoded)
	hash := hmac.New(sha256.New, []byte(h.secret))
	hash.Write(body)
	logrus.Debug(hash)
	return hmac.Equal(decoded, hash.Sum(nil))
}
