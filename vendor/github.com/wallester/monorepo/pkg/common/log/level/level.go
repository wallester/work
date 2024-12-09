package level

import (
	"github.com/rs/zerolog"
)

// Level defines logging level.
type Level int8

const (
	Trace Level = -1
	Debug Level = 0
	Info  Level = 1
	Warn  Level = 2
	Error Level = 3
	Fatal Level = 4
	Panic Level = 5
)

var Levels = []Level{
	Trace,
	Debug,
	Info,
	Warn,
	Error,
	Fatal,
	Panic,
}

func init() {
	// Check that all levels are equal to zerolog levels.
	if Trace != Level(zerolog.TraceLevel) {
		panic("Trace level is not equal to zerolog.TraceLevel")
	}

	if Debug != Level(zerolog.DebugLevel) {
		panic("Debug level is not equal to zerolog.DebugLevel")
	}

	if Info != Level(zerolog.InfoLevel) {
		panic("Info level is not equal to zerolog.InfoLevel")
	}

	if Warn != Level(zerolog.WarnLevel) {
		panic("Warn level is not equal to zerolog.WarnLevel")
	}

	if Error != Level(zerolog.ErrorLevel) {
		panic("Error level is not equal to zerolog.ErrorLevel")
	}

	if Fatal != Level(zerolog.FatalLevel) {
		panic("Fatal level is not equal to zerolog.FatalLevel")
	}

	if Panic != Level(zerolog.PanicLevel) {
		panic("Panic level is not equal to zerolog.PanicLevel")
	}
}
