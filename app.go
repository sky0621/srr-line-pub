package pub

// AppIF ...
type AppIF interface {
	Start(arg *Arg) ExitCode
}

// App ...
type App struct {
}

// NewApp ...
func NewApp(arg *Arg) *App {
	// FIXME config.tomlのパース
	// FIXME loggerのセッティング
	// FIXME AWS接続のセッティング
	// FIXME LINE接続のセッティング
	// FIXME Ctxへの詰め込み
	app := &App{}
	return app
}

// Start ...
func (a *App) Start() ExitCode {

	return ExitCodeOK
}
