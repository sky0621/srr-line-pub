package pub

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/line/line-bot-sdk-go/linebot"
)

type webHandler struct {
	logger      *appLogger
	config      *config
	awsHandler  awsHandlerIF
	lineHandler lineHandlerIF
}

func (h *webHandler) log() *logrus.Entry {
	return h.logger.entry
}

func (h *webHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	h.log().Debug("HandleFunc will start")
	events, err := h.lineHandler.parseRequest(r)
	if err != nil {
		h.log().Errorf("error: %#v", err)
		return
	}
	h.log().Debugf("LINE Messages will handle eventLength:%d", len(events))

	for _, event := range events {
		h.log().Debug(fmt.Sprintf("event: %#v", event))

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				h.log().Debug(fmt.Sprintf("TextMessage: %#v", message))
				var newMsg *linebot.TextMessage
				if "あぶない" == message.Text {
					newMsg = linebot.NewTextMessage("ばしょをちずでおしえて！")
				} else {
					newMsg = linebot.NewTextMessage(message.Text + "!?")
				}
				h.log().Debug(fmt.Sprintf("newMsg %#v", newMsg))

				repMsg := h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg)
				h.log().Debug(fmt.Sprintf("repMsg %#v", repMsg))

				if _, err = h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg).Do(); err != nil {
					h.log().Error("ReplyMessage", err)
					continue
				}

				h.log().Debug("SQS will insert")
				sqsRes, err := h.awsHandler.getSqsHandler().sendMessage(message.Text)
				if err != nil {
					h.log().Error("sqsCli.SendMessage", err)
					continue
				}
				h.log().Debug(fmt.Sprintf("sqsRes %#v", sqsRes))

			case *linebot.LocationMessage:
				h.log().Debug(fmt.Sprintf("LocationMessage %#v", message))
				lat := message.Latitude
				lon := message.Longitude
				addr := message.Address
				retMsg := fmt.Sprintf("じゅうしょは、%s \n緯度：%f\n経度：%f\nだね。ありがとう。みんなにもおしえてあげるね。", addr, lat, lon)
				h.log().Debug(fmt.Sprintf("retMsg %#v", retMsg))
				newMsg := linebot.NewTextMessage(retMsg)
				repMsg := h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg)
				h.log().Debug(fmt.Sprintf("repMsg %#v", repMsg))

				if _, err = h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg).Do(); err != nil {
					h.log().Error("ReplyMessage", err)
					continue
				}

				h.log().Debug("SQS will insert")
				sqsRes, err := h.awsHandler.getSqsHandler().sendMessage(fmt.Sprintf("lat:%v, lon:%v, addr:%v", lat, lon, addr))
				if err != nil {
					h.log().Error("sqsCli.SendMessage", err)
					continue
				}
				h.log().Debug(fmt.Sprintf("sqsRes %#v", sqsRes))

			default:
				h.log().Debug(fmt.Sprintf("Other Message %#v", message))
			}
		}

	}

}
