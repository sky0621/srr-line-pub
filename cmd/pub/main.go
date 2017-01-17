package main

import (
	"flag"
	"log"
	"os"

	pub "github.com/sky0621/srr-line-pub"
)

func main() {
	os.Exit(realMain())
}

func realMain() (exitCode pub.ExitCode) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("Panic occured. ERR: %+v", err)
		}
	}()

	return wrappedMain()
}

func wrappedMain() pub.ExitCode {
	f := flag.String("f", "../config.toml", "Config File Fullpath")
	flag.Parse()

	// Viperグローバル持ち
	err := pub.ReadConfig(f)
	if err != nil {
		panic(err)
	}

	app, exitCode := pub.NewApp(&pub.Credential{
		LineAccessToken:    os.Getenv("LINE_ACCESS_TOKEN"),
		LineChannelSecret:  os.Getenv("LINE_CHANNEL_SECRET"),
		AwsAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}, pub.NewConfig())
	if exitCode > pub.ExitCodeOK {
		return exitCode
	}
	defer app.Close()

	return app.Start()
}
