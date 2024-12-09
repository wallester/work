package log

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	logapi "github.com/wallester/monorepo/pkg/common/log/api"
)

// Entry is a log entry. It can be used to build a log entry and then write it.
// It is not thread safe to use same instance of Entry in multiple goroutines.
// You must call one of the methods to write the log entry: Debug, Error, Fatal, Info, Panic, Warn.
// Failure to do so will discard the log entry.
// Example:
//
//	log.NewEntry().String("foo", "bar").Int("baz", 1).Info("hello")
//
//	// Or:
//	log.NewEntry().Map(log.M{
//		"foo": "bar",
//		"baz": 1,
//	}).Info("hello")
//
// But not:
//
//	log.NewEntry().String("foo", "bar").Int("baz", 1)
//	// This will discard the log entry.
type Entry struct {
	logger logapi.ILogger
	fields map[string]any
}

func (e *Entry) Bytes(name string, value []byte) logapi.IEntry {
	return e.Field(name, value)
}

func (e *Entry) Int(name string, value int) logapi.IEntry {
	return e.Field(name, value)
}

func (e *Entry) Int64(name string, value int64) logapi.IEntry {
	return e.Field(name, value)
}

func (e *Entry) Float64(name string, value float64) logapi.IEntry {
	return e.Field(name, value)
}

func (e *Entry) String(name, value string) logapi.IEntry {
	return e.Field(name, value)
}

func (e *Entry) Any(name string, value any) logapi.IEntry {
	return e.Field(name, value)
}

func (e *Entry) UUID4(name string, value strfmt.UUID4) logapi.IEntry {
	return e.String(name, value.String())
}

func (e *Entry) Strings(name string, value []string) logapi.IEntry {
	return e.String(name, fmt.Sprintf("%v", value))
}

func (e *Entry) Field(name string, value any) logapi.IEntry {
	if e.fields == nil {
		e.fields = make(map[string]any)
	}

	e.fields[name] = value
	return e
}

// M is just a shortcut for map[string]any
type M map[string]any

func (e *Entry) Map(fields map[string]any) logapi.IEntry {
	for k, v := range fields {
		e.Field(k, v)
	}

	return e
}

func (e *Entry) Bool(name string, value bool) logapi.IEntry {
	return e.Field(name, value)
}

// Call any of the following to write the log entry.

func (e *Entry) Debug(msg string) {
	e.logger.Debug(msg, e.fields)
}

func (e *Entry) Error(err error) {
	e.logger.Error(err, e.fields)
}

func (e *Entry) Fatal(err error) {
	e.logger.Fatal(err, e.fields)
}

func (e *Entry) Info(msg string) {
	e.logger.Info(msg, e.fields)
}

func (e *Entry) Panic(err error) {
	e.logger.Panic(err, e.fields)
}

func (e *Entry) Warn(err error) {
	e.logger.Warn(err, e.fields)
}

// Return any of the following to write the log entry and return the error.

func (e *Entry) ReturnError(err error) error {
	e.Error(err)
	return err
}

func (e *Entry) ReturnFatal(err error) error {
	e.Fatal(err)
	return err
}

func (e *Entry) ReturnPanic(err error) error {
	e.Panic(err)
	return err
}

func (e *Entry) ReturnWarn(err error) error {
	e.Warn(err)
	return err
}
