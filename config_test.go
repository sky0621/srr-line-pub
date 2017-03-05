package pub

import "testing"

var expected = &Config{
	server: &serverConfig{
		host: "localhost",
		port: ":8080",
	},
	logger: &loggerConfig{
		appName:  "srr-line-pub",
		filepath: "./srr-line-pub.log",
		level:    "debug",
		server: &serverConfig{
			host: "localhost",
			port: ":8080",
		},
	},
	aws: &awsConfig{
		sqs: &sqsConfig{
			region:   "ap-northeast-1",
			endpoint: "localhost",
			name:     "queuename",
		},
	},
	line: &lineConfig{
		webhookURL: "/local/path",
	},
}

func TestNewConfig(t *testing.T) {
	ReadConfig("./cmd/config.toml")
	actual := NewConfig()
	if expected.String() != actual.String() {
		t.Errorf("\nExpect is \n%#v\nActual is \n%#v", expected.String(), actual.String())
	}
}
