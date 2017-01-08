package pub

// NewArg ...
func NewArg(configFilePath, awsAccessKeyID, awsSecretAccessKey, lineChannelSecret, lineAccessToken string) *Arg {
	a := &Arg{
		configFilePath:     configFilePath,
		awsAccessKeyID:     awsAccessKeyID,
		awsSecretAccessKey: awsSecretAccessKey,
		lineChannelSecret:  lineChannelSecret,
		lineAccessToken:    lineAccessToken,
	}
	return a
}

// Arg ...
type Arg struct {
	configFilePath     string
	awsAccessKeyID     string
	awsSecretAccessKey string
	lineChannelSecret  string
	lineAccessToken    string
}
