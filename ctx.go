package pub

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/line/line-bot-sdk-go/linebot"
)

type ctx struct {
	logger     *AppLogger
	config     *Config
	awsSession *session.Session
	sqsCli     *sqs.SQS
	lineCli    *linebot.Client
}
