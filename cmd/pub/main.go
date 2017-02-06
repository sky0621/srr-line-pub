package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	pub "github.com/sky0621/srr-line-pub"
	"github.com/uber-go/zap"
)

func main() {
	code := realMain()
	zap.Logger.Info("ExitCode", zap.Int("code", 1))
	logrus.Infof("ExitCode: %v", code)
	os.Exit(int(code))
}

func realMain() (exitCode pub.ExitCode) {
	defer func() {
		err := recover()
		if err != nil {
			logrus.Errorf("Panic occured. ERR: %+v", err)
			exitCode = pub.ExitCodePanic
		}
	}()

	return wrappedMain()
}

func wrappedMain() pub.ExitCode {
	credential, config := setup()
	ctx, err := pub.NewCtx(credential, config)
	if err != nil {
		logrus.Errorf("[wrappedMain][call NewCtx()] %#v\n", err)
		return ExitCodeCtxError
	}

	return pub.StartApp(credential, config)
}
