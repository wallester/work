package api

import (
	"github.com/go-openapi/strfmt"
	"github.com/rs/zerolog"
)

type ILogger interface {
	// Bytes etc. are typed helper methods to add fields to the log entry.
	Bytes(name string, value []byte) ILogger
	Int(name string, value int) ILogger
	Int64(name string, value int64) ILogger
	String(name, value string) ILogger
	Any(name string, value any) ILogger
	UUID4(name string, value strfmt.UUID4) ILogger
	Strings(name string, value []string) ILogger
	Field(name string, value any) ILogger
	Map(fields map[string]any) ILogger
	Bool(name string, value bool) ILogger
	// Write implements io.Writer
	Write(p []byte) (int, error)
	// Error etc. are actual logging methods.
	Error(err error, fieldArgs ...map[string]any)
	Fatal(err error, fieldArgs ...map[string]any)
	Panic(err error, fieldArgs ...map[string]any)
	Warn(err error, fieldArgs ...map[string]any)
	Debug(msg string, fieldArgs ...map[string]any)
	Trace(msg string, fieldArgs ...map[string]any)
	Info(msg string, fieldArgs ...map[string]any)
	// NewEntry is used to chain fields and log messages.
	// It does not write the log entry - you need to call one of the methods above to do that.
	// It does not create a new logger instance - it returns the same logger instance.
	NewEntry() IEntry
	ZeroLogger() zerolog.Logger
	// NewChildLogger creates a new child logger.
	NewChildLogger() ILogger
	IsChildLogger() bool
}
