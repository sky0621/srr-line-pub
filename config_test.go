package pub

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	arg := &Arg{
		configFilePath:     "./cmd/pub/config.toml",
		awsAccessKeyID:     "j22frhd0ja4j587y36p9p38o5s2d8rx7",
		awsSecretAccessKey: "ro1cyz1w6mol1kbwkwt8jo7jet5gxz7z",
		lineChannelSecret:  "ld8njsgn42cbtbnkdff8h9ii7jzpsxua",
		lineAccessToken:    "ykx7f1ci90dlart11m7h6uzedtu5ymo0",
	}

	actual := newConfig(arg)

	expected := &Config{
		Arg:     arg,
		AppName: "srr-line-pub",
		Server: &ServerConfig{
			Host: "localhost",
			Port: ":8080",
		},
		Logger: &LoggerConfig{
			Filepath: "srr-line-pub.log",
			Level:    "debug",
		},
		Aws: &AwsConfig{
			Sqs: &SqsConfig{
				Region:   "ap-northeast-1",
				QueueURL: "【SQSのURL】",
			},
		},
		Line: &LineConfig{
			WebhookUrl: "【LINEのWebhook URL】",
		},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nExpect is %#v\nActual is %#v", expected, actual)
	}
}
