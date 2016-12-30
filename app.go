package pub

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/uber-go/zap"
)

// App ...
type App struct {
	ctx *ctx
}

// NewApp ...
func NewApp(arg *Arg) (*App, int) {
	config := newConfig(arg)

	logger, err := newAppLogger(config)
	if err != nil {
		return nil, ExitCodeLogSetupError
	}
	defer logger.logfile.Close()

	// AWS接続のセッティング
	sess, err := session.NewSession(&aws.Config{Region: aws.String(config.Aws.Sqs.Region)})
	if err != nil {
		logger.entry.Error("error: %#v", zap.Error(err))
		return nil, ExitCodeAwsSettingError
	}
	sqsCli := sqs.New(sess)

	logger.entry.Info("AWS connect setting done")

	// LINE接続のセッティング
	lineCli, err := linebot.New(config.Arg.lineChannelSecret, config.Arg.lineAccessToken)
	if err != nil {
		logger.entry.Error("error: %#v", zap.Error(err))
		return nil, ExitCodeLineSettingError
	}
	logger.entry.Info("LINE connect setting done")

	app := &App{
		ctx: &ctx{
			logger:     logger,
			config:     config,
			awsSession: sess,
			sqsCli:     sqsCli,
			lineCli:    lineCli,
		},
	}
	return app, ExitCodeOK
}

// Start ...
func (a *App) Start() int {
	a.ctx.logger.entry.Info("App start")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler := &webHandler{ctx: a.ctx}
	e.POST("/srr/webhook", handler.HandlerFunc)

	err := e.Start(a.ctx.config.Server.Port)
	if err != nil {
		a.ctx.logger.entry.Error("error: %#v", zap.Error(err))
		return ExitCodeServerStartError
	}

	return ExitCodeOK
}
