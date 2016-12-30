package pub

import (
	"os"

	"github.com/Sirupsen/logrus"
)

// AppLogger ...
type AppLogger struct {
	entry   *logrus.Entry
	logfile *os.File
}

// newAppLogger ...
func newAppLogger(c *Config) (*AppLogger, error) {
	_, err := os.Stat(c.Logger.Filepath)
	var logfile *os.File
	if err == nil {
		logfile, err = os.OpenFile(c.Logger.Filepath, os.O_APPEND, 0666)
	} else {
		logfile, err = os.Create(c.Logger.Filepath)
	}
	if err != nil {
		return nil, err
	}

	logrusEntry := logrus.WithFields(logrus.Fields{
		"host":   c.Server.Host,
		"port":   c.Server.Port,
		"system": c.AppName,
	})
	logrusEntry.Logger.Formatter = new(logrus.TextFormatter)
	logrusEntry.Logger.Formatter = new(logrus.JSONFormatter) // default

	return &AppLogger{entry: logrusEntry, logfile: logfile}, nil
}

// Close ...
func (l *AppLogger) Close() {
	l.logfile.Close()
}
