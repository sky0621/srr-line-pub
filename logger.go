package srrlinepub

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
func newLogger(cfg *config) (*logger, error) {
	logrusEntry := logrus.WithFields(logrus.Fields{
		"host":   cfg.Host,
		"port":   cfg.Port,
		"system": "srr-line-pub",
	})
	logrusEntry.Logger.Formatter = new(logrus.TextFormatter)
	logrusEntry.Logger.Formatter = new(logrus.JSONFormatter) // default

	_, err := os.Stat(cfg.Filepath)
	var logfile *os.File
	if err == nil {
		logfile, err = os.OpenFile(cfg.Filepath, os.O_APPEND, 0666)
	} else {
		logfile, err = os.Create(cfg.Filepath)
	}
	if err != nil {
		return nil, err
	}
	logrusEntry.Logger.Out = logfile

	level, err := logrus.ParseLevel(cfg.LogLevel)
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
