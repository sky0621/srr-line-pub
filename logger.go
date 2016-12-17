package pub

import (
	"os"

	"github.com/Sirupsen/logrus"
)

// logger ...
type logger struct {
	entry   *logrus.Entry
	logfile *os.File
}

// newLogger ...
func newLogger(c *Config) (*logger, error) {
	logrusEntry := logrus.WithFields(logrus.Fields{
		"Host":   c.ServerConfig.Host,
		"Port":   c.ServerConfig.Port,
		"system": c.AppName,
	})
	logrusEntry.Logger.Formatter = new(logrus.TextFormatter)
	logrusEntry.Logger.Formatter = new(logrus.JSONFormatter) // default

	_, err := os.Stat(c.LogConfig.Filepath)
	var logfile *os.File
	if err == nil {
		logfile, err = os.OpenFile(c.LogConfig.Filepath, os.O_APPEND, 0666)
	} else {
		logfile, err = os.Create(c.LogConfig.Filepath)
	}
	if err != nil {
		return nil, err
	}
	logrusEntry.Logger.Out = logfile

	level, err := logrus.ParseLevel(c.LogConfig.LogLevel)
	if err != nil {
		return nil, err
	}
	logrusEntry.Logger.Level = level

	return &logger{entry: logrusEntry, logfile: logfile}, nil
}

// close ...
func (l *logger) close() {
	l.logfile.Close()
}
