package pub

// AppIF ...
type AppIF interface {
	Start(arg *Arg) int
}

// App ...
type App struct {
	logger *logger
}

// NewApp ...
func NewApp(arg *Arg) (*App, int) {
	config, err := newConfig(arg)
	if err != nil {
		return nil, ExitCodeConfigError
	}

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
	}
	return app, ExitCodeOK
}

// Start ...
func (a *App) Start() int {

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
