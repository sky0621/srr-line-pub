package pub

const (
	ExitCodeOK int = iota + 1
	ExitCodeArgsError
	ExitCodeLogSetupError
	ExitCodeConfigError
	ExitCodeAwsSettingError
	ExitCodeLineSettingError
	ExitCodeServerStartError
	ExitCodeError
	ExitCodePanic
)

const (
	constEnvLocal = "local"
)
