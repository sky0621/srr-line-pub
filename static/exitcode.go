package static

type ExitCode int

const (
	ExitCodeOK ExitCode = iota + 1
	ExitCodeArgsError
	ExitCodeLogSetupError
	ExitCodeConfigError
	ExitCodeCtxError
	ExitCodeAwsSettingError
	ExitCodeLineSettingError
	ExitCodeServerStartError
	ExitCodeError
	ExitCodePanic
)
