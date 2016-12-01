package pub

// ExitCode ...
type ExitCode int

const (
	ExitCodeOK ExitCode = iota
	ExitCodeArgsError
	ExitCodeLogSetupError
	ExitCodeConfigError
	ExitCodeError
	ExitCodePanic
)
