package hook

import (
	"github.com/juju/errors"
	"github.com/rs/zerolog"
)

// Level attaches "level" field to event as integer.
type Level struct{}

func NewLevel() Level {
	return Level{}
}

// Run executes the hook.
func (h Level) Run(e *zerolog.Event, level zerolog.Level, _ string) {
	if level == zerolog.NoLevel {
		return
	}

	e.Int(levelEvent, levelToInt(level))
}

// private

// levelEvent is a constant which represents int field level in zerolog event.
const levelEvent = "level"

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
	case zerolog.TraceLevel:
		return 8 // Trace: trace messages
	default:
		panic(errors.New("unknown log level: " + level.String()))
	}
}
