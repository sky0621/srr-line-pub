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
		AppName:            "srr-line-pub",
		ServerHost:         "localhost",
		ServerPort:         ":8080",
		LogFilepath:        "srr-line-pub.log",
		LogLevel:           "debug",
		AwsAccessKeyID:     "j22frhd0ja4j587y36p9p38o5s2d8rx7",
		AwsSecretAccessKey: "ro1cyz1w6mol1kbwkwt8jo7jet5gxz7z",
		SqsRegion:          "ap-northeast-1",
		SqsQueueURL:        "https://sqs.ap-northeast-1.amazonaws.com/065886101085/sri-line-message",
		LineChannelSecret:  "ld8njsgn42cbtbnkdff8h9ii7jzpsxua",
		LineAccessToken:    "ykx7f1ci90dlart11m7h6uzedtu5ymo0",
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nExpect is %#v\nActual is %#v", expected, actual)
	}
}
