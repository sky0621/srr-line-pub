package pub

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/line/line-bot-sdk-go/linebot"
)

type webHandler struct {
	config      *Config
	awsHandler  awsHandlerIF
	lineHandler lineHandlerIF
}

func (h *webHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("HandleFunc will start")
	events, err := h.lineHandler.parseRequest(r)
	if err != nil {
		logrus.Errorf("error: %#v", err)
		return
	}
	logrus.Debugf("LINE Messages will handle eventLength:%d", len(events))

	for _, event := range events {
		logrus.Debug(fmt.Sprintf("event: %#v", event))

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				logrus.Debug(fmt.Sprintf("TextMessage: %#v", message))
				var newMsg *linebot.TextMessage
				if "あぶない" == message.Text {
					newMsg = linebot.NewTextMessage("ばしょをちずでおしえて！")
				} else {
					newMsg = linebot.NewTextMessage(message.Text + "!?")
				}
				logrus.Debug(fmt.Sprintf("newMsg %#v", newMsg))

				repMsg := h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg)
				logrus.Debug(fmt.Sprintf("repMsg %#v", repMsg))

				if _, err = h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg).Do(); err != nil {
					logrus.Error("ReplyMessage", err)
					continue
				}

				logrus.Debug("SQS will insert")
				sqsRes, err := h.awsHandler.getSqsHandler().sendMessage(message.Text)
				if err != nil {
					logrus.Error("sqsCli.SendMessage", err)
					continue
				}
				logrus.Debug(fmt.Sprintf("sqsRes %#v", sqsRes))

			case *linebot.LocationMessage:
				logrus.Debug(fmt.Sprintf("LocationMessage %#v", message))
				lat := message.Latitude
				lon := message.Longitude
				addr := message.Address
				retMsg := fmt.Sprintf("じゅうしょは、%s \n緯度：%f\n経度：%f\nだね。ありがとう。みんなにもおしえてあげるね。", addr, lat, lon)
				logrus.Debug(fmt.Sprintf("retMsg %#v", retMsg))
				newMsg := linebot.NewTextMessage(retMsg)
				repMsg := h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg)
				logrus.Debug(fmt.Sprintf("repMsg %#v", repMsg))

				if _, err = h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg).Do(); err != nil {
					logrus.Error("ReplyMessage", err)
					continue
				}

				logrus.Debug("SQS will insert")
				sqsRes, err := h.awsHandler.getSqsHandler().sendMessage(fmt.Sprintf("lat:%v, lon:%v, addr:%v", lat, lon, addr))
				if err != nil {
					logrus.Error("sqsCli.SendMessage", err)
					continue
				}
				logrus.Debug(fmt.Sprintf("sqsRes %#v", sqsRes))

			default:
				logrus.Debug(fmt.Sprintf("Other Message %#v", message))
			}
		}

	}

}
