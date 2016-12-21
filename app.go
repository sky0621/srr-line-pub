package pub

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// AppIF ...
type AppIF interface {
	Start() int
}

// App ...
type App struct {
	logger *logger
	config *Config
}

// NewApp ...
func NewApp(arg *Arg) (AppIF, int) {
	config := newConfig(arg)

	logger, err := newLogger(config)
	if err != nil {
		return nil, ExitCodeLogSetupError
	}
	defer logger.close()

	// FIXME AWS接続のセッティング

	// FIXME LINE接続のセッティング

	// FIXME Ctxへの詰め込み

	app := &App{
		logger: logger,
		config: config,
	}
	return app, ExitCodeOK
}

// Start ...
func (a *App) Start() int {
	sqs.New(&aws.Config{Region: aws.String(a.config.SqsRegion)})

	// e := echo.New()
	// e.Use(middleware.Logger())
	//
	// e.POST("/", func(c echo.Context) error {
	// 	bot, err := linebot.New(&ls, &lt)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	events, err := bot.ParseRequest(c.Request())
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	//
	// 	for _, event := range events {
	// 		// &sqs.
	// 	}
	//
	// 	// FIXME events を処理する。
	// 	// FIXME aws-sdk-go/sqs にてキューに投入
	//
	// 	return c.JSON(http.StatusOK, interface{})
	// })

	return ExitCodeOK
}
