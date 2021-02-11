package logger_test

import (
	"testing"

	l "github.com/joinself/logger"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestSetGlobalLevel(t *testing.T) {
	var tests = []struct {
		name        string
		level       int
		expectation zerolog.Level
	}{
		{
			"default level",
			l.DEFAULT,
			zerolog.InfoLevel,
		},
		{
			"debug level",
			l.DEBUG,
			zerolog.DebugLevel,
		},
		{
			"info level",
			l.INFO,
			zerolog.InfoLevel,
		},
		{
			"notice level",
			l.NOTICE,
			zerolog.InfoLevel,
		},
		{
			"warning level",
			l.WARNING,
			zerolog.WarnLevel,
		},
		{
			"error level",
			l.ERROR,
			zerolog.ErrorLevel,
		},
		{
			"critical level",
			l.CRITICAL,
			zerolog.FatalLevel,
		},
		{
			"alert level",
			l.ALERT,
			zerolog.PanicLevel,
		},
		{
			"emergency level",
			l.EMERGENCY,
			zerolog.PanicLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.SetGlobalLevel(tt.level)
			l := zerolog.GlobalLevel()

			assert.Equal(t, l, tt.expectation)
		})
	}
}
