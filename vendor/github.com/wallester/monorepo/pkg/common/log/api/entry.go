package api

import (
	"github.com/go-openapi/strfmt"
)

type IEntry interface {
	INamedFields
	// Bytes etc. are typed helper methods to add fields to the log entry.
	Bytes(name string, value []byte) IEntry
	Int(name string, value int) IEntry
	Int64(name string, value int64) IEntry
	Float64(name string, value float64) IEntry
	String(name, value string) IEntry
	Any(name string, value any) IEntry
	UUID4(name string, value strfmt.UUID4) IEntry
	Strings(name string, value []string) IEntry
	Field(name string, value any) IEntry
	Map(fields map[string]any) IEntry
	Bool(name string, value bool) IEntry
	// Debug etc. method need to be called to write the log entry.
	Debug(msg string)
	Error(err error)
	Fatal(err error)
	Info(msg string)
	Panic(err error)
	Warn(err error)
	// ReturnError etc. write the log entry and return the error.
	ReturnError(err error) error
	ReturnFatal(err error) error
	ReturnPanic(err error) error
	ReturnWarn(err error) error
}
