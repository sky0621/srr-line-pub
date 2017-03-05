package pub

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sky0621/go-lib/log"
	"github.com/sky0621/srr-line-pub/global"
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
		global.L.Log(log.W, r.Header.Get("X-LINE-Signature"))
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
	global.L.Log(log.D, "##### validateSignature #####")
	global.L.Log(log.D, decoded)
	hash := hmac.New(sha256.New, []byte(h.secret))
	hash.Write(body)
	global.L.Log(log.D, hash)
	return hmac.Equal(decoded, hash.Sum(nil))
}
