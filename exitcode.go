package pub

const (
	ExitCodeOK int = iota + 1
	ExitCodeArgsError
	ExitCodeLogSetupError
	ExitCodeConfigError
	ExitCodeError
	ExitCodePanic
)
