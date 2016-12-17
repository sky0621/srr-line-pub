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

	actual, err := newConfig(arg)
	if err != nil {
		t.Error(err)
	}

	expected := &Config{
		AppName: "srr-line-pub",
		ServerConfig: &ServerConfig{
			Host: "localhost",
			Port: ":8080",
		},
		LogConfig: &LogConfig{
			Filepath: "srr-line-pub.log",
			LogLevel: "debug",
		},
		AwsConfig: &AwsConfig{
			AccessKeyID:     "j22frhd0ja4j587y36p9p38o5s2d8rx7",
			SecretAccessKey: "ro1cyz1w6mol1kbwkwt8jo7jet5gxz7z",
			SqsConfig: &SqsConfig{
				Region:   "ap-northeast-1",
				QueueURL: "https://sqs.ap-northeast-1.amazonaws.com/065886101085/sri-line-message",
			},
		},
		LineConfig: &LineConfig{
			ChannelSecret: "ld8njsgn42cbtbnkdff8h9ii7jzpsxua",
			AccessToken:   "ykx7f1ci90dlart11m7h6uzedtu5ymo0",
		},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nExpect is %v\nActual is %v", expected, actual)
	}
}
