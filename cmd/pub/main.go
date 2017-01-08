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

func realMain() (exitCode int) {
	// treat panic
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("Panic occured. ERR: %+v", err)
			// FIXME 後始末

		}
	}()

	return wrappedMain()
}

func wrappedMain() int {
	arg := parseFlag()
	app, exitCode := pub.NewApp(arg)
	if exitCode > pub.ExitCodeOK {
		return exitCode
	}
	defer app.Close()

	return app.Start()
}

func parseFlag() *pub.Arg {
	f := flag.String("f", "./config.toml", "Config File Fullpath")
	// LINE-API
	ls := flag.String("ls", "channelSecret", "LINE ChannelSecret")
	lt := flag.String("lt", "accessToken", "LINE AccessToken")
	flag.Parse()
	arg := pub.NewArg(*f, *ls, *lt)
	return arg
}
