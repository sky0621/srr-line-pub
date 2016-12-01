package pub

// AppIF ...
type AppIF interface {
	Start(arg *Arg) ExitCode
}

// App ...
type App struct {
}

// Start ...
func (a *App) Start(arg *Arg) ExitCode {

	return ExitCodeOK
}
