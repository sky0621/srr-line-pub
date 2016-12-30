package pub

import (
	"os"
	"time"

	"github.com/uber-go/zap"
)

// AppLogger ...
type AppLogger struct {
	entry   zap.Logger
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

	logger := zap.New(
		zap.NewJSONEncoder(jstTimeFormatter("timestamp")),
		zap.Output(logfile),
	)
	logger.With(
		zap.String("Host", c.Server.Host),
		zap.String("Port", c.Server.Port),
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

// Close ...
func (l *AppLogger) Close() {
	l.logfile.Close()
}
