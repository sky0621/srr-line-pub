package pub

import (
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/uber-go/zap"
)

// AppLogger ...
type AppLogger struct {
	entry   zap.Logger
	logfile *os.File
}

// AppLogger ...
type AppLogger2 struct {
	entry   *logrus.Entry
	logfile *os.File
}

// newAppLogger ...
func newAppLogger(c *Config) (*AppLogger, error) {
	_, err := os.Stat(c.LogFilepath)
	var logfile *os.File
	if err == nil {
		logfile, err = os.OpenFile(c.LogFilepath, os.O_APPEND, 0666)
	} else {
		logfile, err = os.Create(c.LogFilepath)
	}
	if err != nil {
		return nil, err
	}

	logger := zap.New(
		zap.NewJSONEncoder(jstTimeFormatter("timestamp")),
		zap.Output(logfile),
	)
	logger.With(
		zap.String("Host", c.ServerHost),
		zap.String("Port", c.ServerPort),
		zap.String("System", c.AppName),
	)
	return &AppLogger{entry: logger, logfile: logfile}, nil
}

func jstTimeFormatter(key string) zap.TimeFormatter {
	return zap.TimeFormatter(func(t time.Time) zap.Field {
		const layout = "2006-01-02 15:04:05"
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		return zap.String(key, t.In(jst).Format(layout))
	})
}

// newAppLogger ...
func newAppLogger2(c *Config) (*AppLogger2, error) {
	logrusEntry := logrus.WithFields(logrus.Fields{
		"Host":   c.ServerHost,
		"Port":   c.ServerPort,
		"system": c.AppName,
	})
	logrusEntry.Logger.Formatter = new(logrus.JSONFormatter)

	_, err := os.Stat(c.LogFilepath)
	var logfile *os.File
	if err == nil {
		logfile, err = os.OpenFile(c.LogFilepath, os.O_APPEND, 0666)
	} else {
		logfile, err = os.Create(c.LogFilepath)
	}
	if err != nil {
		return nil, err
	}
	logrusEntry.Logger.Out = logfile

	level, err := logrus.ParseLevel(c.LogLevel)
	if err != nil {
		return nil, err
	}
	logrusEntry.Logger.Level = level

	return &AppLogger2{entry: logrusEntry, logfile: logfile}, nil
}

// Close ...
func (l *AppLogger) Close() {
	l.logfile.Close()
}
