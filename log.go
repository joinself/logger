package logger

import (
	"context"

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

// Message represents a log message
type Message struct {
	e *zerolog.Event
}

// Msg outputs a message
func (m *Message) Msg(msg string) {
	m.e.Msg(msg)
}

// Context adds values from a context to the log message
func (m *Message) Context(ctx context.Context) *Message {
	sessionID := ctx.Value("session_id")
	if sessionID != nil {
		m.e.Str("session_id", sessionID.(string))
	}

	selfID := ctx.Value("self_id")
	if selfID != nil {
		m.e.Str("self_id", selfID.(string))
	}

	return m
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
func Debug() *Message {
	return &Message{zl.Debug().Int("severity", DEBUG)}
}

// Info routine information, such as ongoing status or performance.
func Info() *Message {
	return &Message{zl.Info().Int("severity", INFO)}
}

// Warn events might cause problems.
func Warn() *Message {
	return &Message{zl.Warn().Int("severity", WARNING)}
}

// Warning events might cause problems.
func Warning() *Message {
	return Warn()
}

// Error events are likely to cause problems.
func Error() *Message {
	return &Message{zl.Error().Int("severity", ERROR)}
}

// Fatal critical events cause more severe problems or outages.
func Fatal() *Message {
	return &Message{zl.Fatal().Int("severity", CRITICAL)}
}

// Panic a person must take an action immediately.
func Panic() *Message {
	return &Message{zl.Panic().Int("severity", ALERT)}
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
