package pub

type ExitCode int

const (
	ExitCodeOK ExitCode = iota + 1
	ExitCodeArgsError
	ExitCodeLogSetupError
	ExitCodeConfigError
	ExitCodeAwsSettingError
	ExitCodeLineSettingError
	ExitCodeServerStartError
	ExitCodeError
	ExitCodePanic
)
