package pub

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/uber-go/zap"
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

func (h *webHandler) HandlerFunc(c echo.Context) error {
	events, err := h.ctx.lineCli.ParseRequest(c.Request())
	if err != nil {
		h.ctx.logger.entry.Error("error: %#v", zap.Error(err))
		return err
	}
	h.ctx.logger.entry.Debug("LINE Messages will handle", zap.Int("eventLength", len(events)))

	for _, event := range events {
		h.ctx.logger.entry.Debug(fmt.Sprintf("event: %#v", event))

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				h.ctx.logger.entry.Debug(fmt.Sprintf("message: %#v", message))
				var newMsg *linebot.TextMessage
				if "あぶない" == message.Text {
					newMsg = linebot.NewTextMessage("ばしょをちずでおしえて！")
				} else {
					newMsg = linebot.NewTextMessage(message.Text + "!?")
				}
				h.ctx.logger.entry.Debug(fmt.Sprintf("newMsg %#v", newMsg))

				repMsg := h.ctx.lineCli.ReplyMessage(event.ReplyToken, newMsg)
				h.ctx.logger.entry.Debug(fmt.Sprintf("repMsg %#v", repMsg))

				if _, err = h.ctx.lineCli.ReplyMessage(event.ReplyToken, newMsg).Do(); err != nil {
					h.ctx.logger.entry.Error("ReplyMessage", zap.Error(err))
					continue
				}

				h.ctx.logger.entry.Debug("SQS will insert")
				sqsParam := &sqs.SendMessageInput{
					QueueUrl:    aws.String(h.ctx.config.Aws.Sqs.QueueURL),
					MessageBody: aws.String(message.Text),
				}
				sqsRes, err := h.ctx.sqsCli.SendMessage(sqsParam)
				if err != nil {
					h.ctx.logger.entry.Error("sqsCli.SendMessage", zap.Error(err))
					continue
				}
				h.ctx.logger.entry.Debug(fmt.Sprintf("sqsRes %#v", sqsRes))

			case *linebot.LocationMessage:
				h.ctx.logger.entry.Debug(fmt.Sprintf("message %#v", message))
				lat := message.Latitude
				lon := message.Longitude
				addr := message.Address
				retMsg := fmt.Sprintf("じゅうしょは、%s \n緯度：%f\n経度：%f\nだね。ありがとう。みんなにもおしえてあげるね。", addr, lat, lon)
				h.ctx.logger.entry.Debug(fmt.Sprintf("retMsg %#v", retMsg))
				newMsg := linebot.NewTextMessage(retMsg)
				repMsg := h.ctx.lineCli.ReplyMessage(event.ReplyToken, newMsg)
				h.ctx.logger.entry.Debug(fmt.Sprintf("repMsg %#v", repMsg))

				if _, err = h.ctx.lineCli.ReplyMessage(event.ReplyToken, newMsg).Do(); err != nil {
					h.ctx.logger.entry.Error("ReplyMessage", zap.Error(err))
					continue
				}

				h.ctx.logger.entry.Debug("SQS will insert")
				sqsParam := &sqs.SendMessageInput{
					QueueUrl:    aws.String(h.ctx.config.Aws.Sqs.QueueURL),
					MessageBody: aws.String(fmt.Sprintf("lat:%v, lon:%v, addr:%v", lat, lon, addr)),
				}
				sqsRes, err := h.ctx.sqsCli.SendMessage(sqsParam)
				if err != nil {
					h.ctx.logger.entry.Error("sqsCli.SendMessage", zap.Error(err))
					continue
				}
				h.ctx.logger.entry.Debug(fmt.Sprintf("sqsRes %#v", sqsRes))

			}
		}

	}

	return nil
}
