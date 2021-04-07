package logger

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"
)

const (
	// DebugLevel defines debug log level.
	DebugLevel int = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled
)

// Message represents a log message
type Message struct {
	e *zerolog.Event
}

// Msg outputs a message
func (m *Message) Msg(format string, args ...interface{}) {
	m.e.Msg(fmt.Sprintf(format, args...))
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
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.LevelFieldName = "severity"
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		// mapping to Cloud Logging LogSeverity
		// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity
		switch l {
		case zerolog.TraceLevel:
			return "DEFAULT"
		case zerolog.DebugLevel:
			return "DEBUG"
		case zerolog.InfoLevel:
			return "INFO"
		case zerolog.WarnLevel:
			return "WARNING"
		case zerolog.ErrorLevel:
			return "ERROR"
		case zerolog.FatalLevel:
			return "CRITICAL"
		case zerolog.PanicLevel:
			return "ALERT"
		case zerolog.NoLevel:
			return "DEFAULT"
		default:
			return "DEFAULT"
		}
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if os.Getenv("LOG_LEVEL") != "" {
		level, err := zerolog.ParseLevel(strings.ToUpper(os.Getenv("LOG_LEVEL")))
		if err != nil {
			Warn().Msg("invalid log level (" + os.Getenv("LOG_LEVEL") + "), defaulting to [INFO]")
		} else {
			zerolog.SetGlobalLevel(level)
		}
	}

}

// SetGlobalLevel sets the global override for log level. If this
// values is raised, all Loggers will use at least this value.
func SetGlobalLevel(level int) {
	zerolog.SetGlobalLevel(zerolog.Level(level))
}

// Debug or trace information.
func Debug() *Message {
	return &Message{zl.Debug()}
}

// Info routine information, such as ongoing status or performance.
func Info() *Message {
	return &Message{zl.Info()}
}

// Warn events might cause problems.
func Warn() *Message {
	return &Message{zl.Warn()}
}

// Warning events might cause problems.
func Warning() *Message {
	return Warn()
}

// Error events are likely to cause problems.
func Error() *Message {
	return &Message{zl.Error()}
}

// Fatal critical events cause more severe problems or outages.
func Fatal() *Message {
	return &Message{zl.Fatal()}
}

// Panic a person must take an action immediately.
func Panic() *Message {
	return &Message{zl.Panic()}
}
