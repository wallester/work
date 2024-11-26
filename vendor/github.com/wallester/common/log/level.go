package log

import (
	"github.com/rs/zerolog"
)

// Level defines logging level.
type Level zerolog.Level

var (
	// DebugLevel defines debug log level.
	DebugLevel = Level(zerolog.DebugLevel)
	// InfoLevel defines info log level.
	InfoLevel = Level(zerolog.InfoLevel)
	// WarnLevel defines warn log level.
	WarnLevel = Level(zerolog.WarnLevel)
	// ErrorLevel defines error log level.
	ErrorLevel = Level(zerolog.ErrorLevel)
	// FatalLevel defines fatal log level.
	FatalLevel = Level(zerolog.FatalLevel)
	// PanicLevel defines panic log level.
	PanicLevel = Level(zerolog.PanicLevel)
)
