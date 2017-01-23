package main

import (
	"flag"
	"os"

	pub "github.com/sky0621/srr-line-pub"
)

func setup() (*pub.Credential, *pub.Config) {
	f := flag.String("f", "../config.toml", "Config File Fullpath")
	flag.Parse()

	// Viperグローバル持ち
	err := pub.ReadConfig(*f)
	if err != nil {
		panic(err)
	}

	return &pub.Credential{
		LineAccessToken:    os.Getenv("LINE_ACCESS_TOKEN"),
		LineChannelSecret:  os.Getenv("LINE_CHANNEL_SECRET"),
		AwsAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}, pub.NewConfig()
}
