package pub

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/line/line-bot-sdk-go/linebot"
)

func webSetup(cfg *loggerConfig) *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	switch cfg.level {
	case "debug":
		e.Logger.SetLevel(log.DEBUG)
	case "info":
		e.Logger.SetLevel(log.INFO)
	case "warn":
		e.Logger.SetLevel(log.WARN)
	case "error":
		e.Logger.SetLevel(log.ERROR)
	}
	logfile, err := logfile(cfg.filepath)
	if err != nil {
		return nil
	}
	e.Logger.SetOutput(logfile)
	return e
}

type webHandler struct {
	logger      *appLogger
	config      *config
	awsHandler  awsHandlerIF
	lineHandler lineHandlerIF
}

// func (h *webHandler) log() *logrus.Entry {
// 	return h.log()
// }

func (h *webHandler) HandlerFunc(c echo.Context) error {
	// h.log().Debug("HandleFunc will start")
	c.Logger().Info("[echo]HandleFunc will start")
	events, err := h.lineHandler.getClient().ParseRequest(c.Request())
	if err != nil {
		// h.log().Errorf("error: %#v", err)
		c.Logger().Errorf("error: %#v", err)
		return err
	}
	// h.log().Debugf("LINE Messages will handle eventLength:%d", len(events))
	c.Logger().Infof("[echo]LINE Messages will handle eventLength:%d", len(events))

	for _, event := range events {
		// h.log().Debug(fmt.Sprintf("event: %#v", event))
		c.Logger().Info(fmt.Sprintf("[echo]event: %#v", event))

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// h.log().Debug(fmt.Sprintf("TextMessage: %#v", message))
				c.Logger().Info(fmt.Sprintf("TextMessage: %#v", message))
				var newMsg *linebot.TextMessage
				if "あぶない" == message.Text {
					newMsg = linebot.NewTextMessage("ばしょをちずでおしえて！")
				} else {
					newMsg = linebot.NewTextMessage(message.Text + "!?")
				}
				// h.log().Debug(fmt.Sprintf("newMsg %#v", newMsg))
				c.Logger().Info(fmt.Sprintf("newMsg %#v", newMsg))

				repMsg := h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg)
				// h.log().Debug(fmt.Sprintf("repMsg %#v", repMsg))
				c.Logger().Info(fmt.Sprintf("repMsg %#v", repMsg))

				if _, err = h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg).Do(); err != nil {
					// h.log().Error("ReplyMessage", err)
					c.Logger().Error("ReplyMessage", err)
					continue
				}

				// h.log().Debug("SQS will insert")
				c.Logger().Info("SQS will insert")
				sqsParam := &sqs.SendMessageInput{
					QueueUrl:    aws.String(h.config.aws.sqs.queueURL),
					MessageBody: aws.String(message.Text),
				}
				sqsRes, err := h.awsHandler.getSqsClient().SendMessage(sqsParam)
				if err != nil {
					// h.log().Error("sqsCli.SendMessage", err)
					c.Logger().Error("sqsCli.SendMessage", err)
					continue
				}
				// h.log().Debug(fmt.Sprintf("sqsRes %#v", sqsRes))
				c.Logger().Info(fmt.Sprintf("sqsRes %#v", sqsRes))

			case *linebot.LocationMessage:
				// h.log().Debug(fmt.Sprintf("LocationMessage %#v", message))
				c.Logger().Info(fmt.Sprintf("LocationMessage %#v", message))
				lat := message.Latitude
				lon := message.Longitude
				addr := message.Address
				retMsg := fmt.Sprintf("じゅうしょは、%s \n緯度：%f\n経度：%f\nだね。ありがとう。みんなにもおしえてあげるね。", addr, lat, lon)
				// h.log().Debug(fmt.Sprintf("retMsg %#v", retMsg))
				c.Logger().Info(fmt.Sprintf("retMsg %#v", retMsg))
				newMsg := linebot.NewTextMessage(retMsg)
				repMsg := h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg)
				// h.log().Debug(fmt.Sprintf("repMsg %#v", repMsg))
				c.Logger().Info(fmt.Sprintf("repMsg %#v", repMsg))

				if _, err = h.lineHandler.getClient().ReplyMessage(event.ReplyToken, newMsg).Do(); err != nil {
					// h.log().Error("ReplyMessage", err)
					c.Logger().Error("ReplyMessage", err)
					continue
				}

				// h.log().Debug("SQS will insert")
				c.Logger().Info("SQS will insert")
				sqsParam := &sqs.SendMessageInput{
					QueueUrl:    aws.String(h.config.aws.sqs.queueURL),
					MessageBody: aws.String(fmt.Sprintf("lat:%v, lon:%v, addr:%v", lat, lon, addr)),
				}
				sqsRes, err := h.awsHandler.getSqsClient().SendMessage(sqsParam)
				if err != nil {
					// h.log().Error("sqsCli.SendMessage", err)
					c.Logger().Error("sqsCli.SendMessage", err)
					continue
				}
				// h.log().Debug(fmt.Sprintf("sqsRes %#v", sqsRes))
				c.Logger().Info(fmt.Sprintf("sqsRes %#v", sqsRes))

			default:
				// h.log().Debug(fmt.Sprintf("Other Message %#v", message))
				c.Logger().Info(fmt.Sprintf("Other Message %#v", message))
			}

		}

	}

	return nil
}
