package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	pub "github.com/sky0621/srr-line-pub"
	"github.com/sky0621/srr-line-pub/static"
	"github.com/uber-go/zap"
)

func main() {
	code := realMain()
	zap.Logger.Info("ExitCode", zap.Int("code", 1))
	logrus.Infof("ExitCode: %v", code)
	os.Exit(int(code))
}

func realMain() (exitCode static.ExitCode) {
	defer func() {
		err := recover()
		if err != nil {
			logrus.Errorf("Panic occured. ERR: %+v", err)
			exitCode = static.ExitCodePanic
		}
	}()

	return wrappedMain()
}

func wrappedMain() static.ExitCode {
	credential, config := setup()
	ctx, err := pub.NewCtx(credential, config)
	if err != nil {
		logrus.Errorf("[wrappedMain][call NewCtx()] %#v\n", err)
		return static.ExitCodeCtxError
	}

	return pub.StartApp(ctx)
}
