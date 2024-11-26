package log

import (
	"github.com/juju/errors"
	"github.com/rs/zerolog"
)

// LevelHook attaches "level" field to event as integer.
type LevelHook struct{}

// Run executes the hook.
func (h LevelHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level != zerolog.NoLevel {
		e.Int("level", levelToInt(level))
	}
}

// private

func levelToInt(level zerolog.Level) int {
	switch level {
	case zerolog.FatalLevel:
		return 0 // Emergency: system is unusable
	case zerolog.PanicLevel:
		return 2 // Critical: critical conditions
	case zerolog.ErrorLevel:
		return 3 // Error: error conditions
	case zerolog.WarnLevel:
		return 4 // Warning: warning conditions
	case zerolog.InfoLevel:
		return 6 // Informational: informational messages
	case zerolog.DebugLevel:
		return 7 // Debug: debug-level messages
	default:
		panic(errors.New("unknown log level: " + level.String()))
	}
}
