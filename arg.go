package pub

// NewArg ...
func NewArg(lineChannelSecret string, lineAccessToken string, awsAccessKeyID string, awsSecretAccessKey string) *Arg {
	a := &Arg{
		LineChannelSecret:  lineChannelSecret,
		LineAccessToken:    lineAccessToken,
		AwsAccessKeyID:     awsAccessKeyID,
		AwsSecretAccessKey: awsSecretAccessKey,
	}
	return a
}

// Arg ...
type Arg struct {
	LineChannelSecret  string
	LineAccessToken    string
	AwsAccessKeyID     string
	AwsSecretAccessKey string
}
