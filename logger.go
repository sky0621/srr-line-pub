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
	logrusEntry := logrus.WithFields(logrus.Fields{
		"host":   c.Server.Host,
		"port":   c.Server.Port,
		"system": c.AppName,
	})
	logrusEntry.Logger.Formatter = new(logrus.JSONFormatter)

	logfile, err := logfile(c.Logger.Filepath)
	if err != nil {
		return nil, err
	}
	logrusEntry.Logger.Out = logfile

	level, err := logrus.ParseLevel(c.Logger.Level)
	if err != nil {
		return nil, err
	}
	logrusEntry.Logger.Level = level

	return &AppLogger{entry: logrusEntry, logfile: logfile}, nil
}

// Close ...
func (l *AppLogger) Close() {
	l.logfile.Close()
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
