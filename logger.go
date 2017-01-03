package pub

import (
	"os"

	"github.com/Sirupsen/logrus"
)

type appLogger struct {
	entry   *logrus.Entry
	logfile *os.File
}

func newAppLogger(c *loggerConfig) (*appLogger, error) {
	logrusEntry := logrus.WithFields(logrus.Fields{
		"host":   c.server.host,
		"port":   c.server.port,
		"system": c.appName,
	})
	logrusEntry.Logger.Formatter = new(logrus.JSONFormatter)

	logfile, err := logfile(c.filepath)
	if err != nil {
		return nil, err
	}
	logrusEntry.Logger.Out = logfile

	level, err := logrus.ParseLevel(c.level)
	if err != nil {
		return nil, err
	}
	logrusEntry.Logger.Level = level

	return &appLogger{entry: logrusEntry, logfile: logfile}, nil
}

// Close ...
func (l *appLogger) Close() error {
	return l.logfile.Close()
}

func logfile(filepath string) (*os.File, error) {
	_, err := os.Stat(filepath)
	var logfile *os.File
	if err == nil {
		logfile, err = os.OpenFile(filepath, os.O_APPEND, 0666)
	} else {
		logfile, err = os.Create(filepath)
	}
	if err != nil {
		return nil, err
	}
	return logfile, nil
}
