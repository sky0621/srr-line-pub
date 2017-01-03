package pub

import (
	"reflect"
	"testing"
)

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
			region:   "ap-northeast-1",
			queueURL: "localhost",
		},
	},
	line: &lineConfig{
		environment: "local",
		webhookURL:  "localhost",
	},
}

func TestNewConfig(t *testing.T) {
	readConfig("./cmd/pub/config.toml")
	actual := newConfig()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nExpect is %#v\nActual is %#v", expected, actual)
	}
}
