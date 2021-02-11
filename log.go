package logger

import (
	"fmt"

	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"
)

const (
	// Log levels on GCP (https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#logseverity)
	// DEFAULT logging level.
	DEFAULT = 0
	// DEBUG logging level.
	DEBUG = 100
	// INFO logging level.
	INFO = 200
	// NOTICE logging level.
	NOTICE = 300
	// WARNING logging level.
	WARNING = 400
	// ERROR logging level.
	ERROR = 500
	// CRITICAL logging level.
	CRITICAL = 600
	// ALERT logging level.
	ALERT = 700
	// EMERGENCY logging level.
	EMERGENCY = 800
)

// L log line.
type L struct {
	Body      string
	SessionID string
	SelfID    string
}

// New creates a new Logger.
func init() {
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

// SetGlobalLevel sets the global override for log level. If this
// values is raised, all Loggers will use at least this value.
func SetGlobalLevel(level int) {
	l := parseLevel(level)
	zerolog.SetGlobalLevel(l)
}

// Debug or trace information.
func Debug(l L) {
	logEvent(DEBUG, l)
}

// Info routine information, such as ongoing status or performance.
func Info(l L) {
	logEvent(INFO, l)
}

// Warn events might cause problems.
func Warn(l L) {
	logEvent(WARNING, l)
}

// Warning events might cause problems.
func Warning(l L) {
	Warn(l)
}

// Error events are likely to cause problems.
func Error(l L) {
	logEvent(ERROR, l)
}

// Fatal critical events cause more severe problems or outages.
func Fatal(l L) {
	logEvent(CRITICAL, l)
}

// Panic a person must take an action immediately.
func Panic(l L) {
	logEvent(ALERT, l)
}

func logEvent(level int, l L) {
	ev := buildEvent(level)
	if ev == nil {
		return
	}
	ev.Int("severity", level)

	if l.SessionID != "" {
		ev.Str("session_id", fmt.Sprint(l.SessionID))
	}

	if l.SelfID != "" {
		ev.Str("self_id", fmt.Sprint(l.SelfID))
	}

	ev.Msg(l.Body)
}

func buildEvent(level int) *zerolog.Event {
	switch level {
	case DEFAULT:
		return zl.Info()
	case DEBUG:
		return zl.Debug()
	case INFO:
		return zl.Info()
	case NOTICE:
		return zl.Info()
	case WARNING:
		return zl.Warn()
	case ERROR:
		return zl.Error()
	case CRITICAL:
		return zl.Fatal()
	case ALERT:
		return zl.Panic()
	case EMERGENCY:
		return zl.Panic()
	default:
		return zl.Info()
	}
}

func parseLevel(level int) zerolog.Level {
	switch level {
	case DEFAULT:
		return zerolog.InfoLevel
	case DEBUG:
		return zerolog.DebugLevel
	case INFO:
		return zerolog.InfoLevel
	case NOTICE:
		return zerolog.InfoLevel
	case WARNING:
		return zerolog.WarnLevel
	case ERROR:
		return zerolog.ErrorLevel
	case CRITICAL:
		return zerolog.FatalLevel
	case ALERT:
		return zerolog.PanicLevel
	case EMERGENCY:
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}

}
