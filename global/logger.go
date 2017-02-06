package global

import (
	"time"

	"github.com/uber-go/zap"
)

const (
	D = iota + 1
	I
	W
	E
	F
	P
)

// Logger ...
type Logger struct {
}

func NewLogger() *Logger {
	logger := zap.New(
		zap.NewJSONEncoder(JSTTimeFormatter("timestamp")),
	)
	return nil
}

// JSTTimeFormatter ...
func JSTTimeFormatter(key string) zap.TimeFormatter {
	return zap.TimeFormatter(func(t time.Time) zap.Field {
		const layout = "2006-01-02 15:04:05"
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		return zap.String(key, t.In(jst).Format(layout))
	})
}
