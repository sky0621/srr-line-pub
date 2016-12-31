package pub

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
)

func webSetup() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}

type webHandler struct {
	ctx *ctx
}

func (h *webHandler) log() *logrus.Entry {
	return h.log()
}

func (h *webHandler) HandlerFunc(c echo.Context) error {
	h.log().Debug("echo HandleFunc will start")
	events, err := h.ctx.lineCli.ParseRequest(c.Request())
	if err != nil {
		h.log().Errorf("error: %#v", err)
		return err
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

				repMsg := h.ctx.lineCli.ReplyMessage(event.ReplyToken, newMsg)
				h.log().Debug(fmt.Sprintf("repMsg %#v", repMsg))

				if _, err = h.ctx.lineCli.ReplyMessage(event.ReplyToken, newMsg).Do(); err != nil {
					h.log().Error("ReplyMessage", err)
					continue
				}

				h.log().Debug("SQS will insert")
				sqsParam := &sqs.SendMessageInput{
					QueueUrl:    aws.String(h.ctx.config.Aws.Sqs.QueueURL),
					MessageBody: aws.String(message.Text),
				}
				sqsRes, err := h.ctx.sqsCli.SendMessage(sqsParam)
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
				repMsg := h.ctx.lineCli.ReplyMessage(event.ReplyToken, newMsg)
				h.log().Debug(fmt.Sprintf("repMsg %#v", repMsg))

				if _, err = h.ctx.lineCli.ReplyMessage(event.ReplyToken, newMsg).Do(); err != nil {
					h.log().Error("ReplyMessage", err)
					continue
				}

				h.log().Debug("SQS will insert")
				sqsParam := &sqs.SendMessageInput{
					QueueUrl:    aws.String(h.ctx.config.Aws.Sqs.QueueURL),
					MessageBody: aws.String(fmt.Sprintf("lat:%v, lon:%v, addr:%v", lat, lon, addr)),
				}
				sqsRes, err := h.ctx.sqsCli.SendMessage(sqsParam)
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

	return nil
}
