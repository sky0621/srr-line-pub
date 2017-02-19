package pub

import (
	"fmt"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sky0621/go-lib/log"
	"github.com/sky0621/srr-line-pub/global"
)

type webHandler struct {
	awsHandler  awsHandlerIF
	lineHandler lineHandlerIF
}

func (h *webHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	global.L.Log(log.D, "HandleFunc will start")
	events, err := h.lineHandler.parseRequest(r)
	if err != nil {
		global.L.Logf(log.E, "error: %#v", err)
		return
	}
	global.L.Logf(log.D, "LINE Messages will handle eventLength:%d", len(events))

	for _, event := range events {
		global.L.Log(log.D, fmt.Sprintf("event: %#v", event))

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				global.L.Log(log.D, fmt.Sprintf("TextMessage: %#v", message))
				var newMsg *linebot.TextMessage
				if "あぶない" == message.Text {
					newMsg = linebot.NewTextMessage("ばしょをちずでおしえて！")
				} else {
					newMsg = linebot.NewTextMessage(message.Text + "!?")
				}
				global.L.Log(log.D, fmt.Sprintf("newMsg %#v", newMsg))

				repMsg := h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg)
				global.L.Log(log.D, fmt.Sprintf("repMsg %#v", repMsg))

				if _, err = h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg).Do(); err != nil {
					global.L.Log(log.E, "ReplyMessage", err)
					continue
				}

				global.L.Log(log.D, "SQS will insert")
				sqsRes, err := h.awsHandler.getSqsHandler().sendMessage(message.Text)
				if err != nil {
					global.L.Log(log.E, "sqsCli.SendMessage", err)
					continue
				}
				global.L.Log(log.D, fmt.Sprintf("sqsRes %#v", sqsRes))

			case *linebot.LocationMessage:
				global.L.Log(log.D, fmt.Sprintf("LocationMessage %#v", message))
				lat := message.Latitude
				lon := message.Longitude
				addr := message.Address
				retMsg := fmt.Sprintf("じゅうしょは、%s \n緯度：%f\n経度：%f\nだね。ありがとう。みんなにもおしえてあげるね。", addr, lat, lon)
				global.L.Log(log.D, fmt.Sprintf("retMsg %#v", retMsg))
				newMsg := linebot.NewTextMessage(retMsg)
				repMsg := h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg)
				global.L.Log(log.D, fmt.Sprintf("repMsg %#v", repMsg))

				if _, err = h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg).Do(); err != nil {
					global.L.Log(log.E, "ReplyMessage", err)
					continue
				}

				global.L.Log(log.D, "SQS will insert")
				sqsRes, err := h.awsHandler.getSqsHandler().sendMessage(fmt.Sprintf("lat:%v, lon:%v, addr:%v", lat, lon, addr))
				if err != nil {
					global.L.Log(log.E, "sqsCli.SendMessage", err)
					continue
				}
				global.L.Log(log.D, fmt.Sprintf("sqsRes %#v", sqsRes))

			default:
				global.L.Log(log.D, fmt.Sprintf("Other Message %#v", message))
			}
		}

	}

}
