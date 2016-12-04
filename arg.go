package pub

// NewArg ...
func NewArg(configFilePath, awsAccessKeyID, awsSecretAccessKey, lineChannelSecret, lineAccessToken string) *Arg {
	a := &Arg{
		ConfigFilePath:     configFilePath,
		AwsAccessKeyID:     awsAccessKeyID,
		AwsSecretAccessKey: awsSecretAccessKey,
		LineChannelSecret:  lineChannelSecret,
		LineAccessToken:    lineAccessToken,
	}
	return a
}

// Arg ...
type Arg struct {
	ConfigFilePath     string
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	LineChannelSecret  string
	LineAccessToken    string
}
