package hook

import (
	"github.com/rs/zerolog"
)

type LoggerName struct {
	name string
}

func NewLoggerName(name string) LoggerName {
	return LoggerName{
		name: name,
	}
}

func (h LoggerName) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	e.Str(loggerName, h.name)
}

// private

const loggerName = "logger-name"
