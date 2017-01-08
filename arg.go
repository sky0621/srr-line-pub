package pub

// NewArg ...
func NewArg(configFilePath, lineChannelSecret, lineAccessToken string) *Arg {
	a := &Arg{
		configFilePath:    configFilePath,
		lineChannelSecret: lineChannelSecret,
		lineAccessToken:   lineAccessToken,
	}
	return a
}

// Arg ...
type Arg struct {
	configFilePath    string
	lineChannelSecret string
	lineAccessToken   string
}
