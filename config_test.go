package pub

import "testing"

var expected = &config{
	environment: "local",
	appName:     "srr-line-pub",
	server: &serverConfig{
		host: "localhost",
		port: ":8080",
	},
	logger: &loggerConfig{
		environment: "local",
		appName:     "srr-line-pub",
		filepath:    "./srr-line-pub.log",
		level:       "debug",
		server: &serverConfig{
			host: "localhost",
			port: ":8080",
		},
	},
	aws: &awsConfig{
		environment: "local",
		sqs: &sqsConfig{
			environment: "local",
			region:      "ap-northeast-1",
			endpoint:    "localhost",
			name:        "queuename",
		},
	},
	line: &lineConfig{
		environment: "local",
		webhookURL:  "/local/path",
	},
}

func TestNewConfig(t *testing.T) {
	readConfig("./cmd/pub/config.toml")
	actual := newConfig()
	if expected.String() != actual.String() {
		t.Errorf("\nExpect is %#v\nActual is %#v", expected.String(), actual.String())
	}
}
