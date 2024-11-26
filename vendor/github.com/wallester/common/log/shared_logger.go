package log

import (
	"os"

	"github.com/juju/errors"
	rollbar "github.com/rollbar/rollbar-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// "So if your writer is a file writer in O_APPEND mode, on most OS it's
// gonna be thread safe as write syscall is atomic in this case. This
// applies to stdout/err if redirected into a file." -- https://github.com/rs/zerolog/issues/29
var sharedLogger *Logger

func init() {
	sharedLogger = &Logger{
		zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

// ConfigureSharedLogger configures shared logger instance and returns a func to close it.
func ConfigureSharedLogger(cfg Configuration) func() {
	logger, closer := New(cfg)
	sharedLogger = logger
	return closer
}

// Panic logs an error and panics.
func Panic(err error, fieldArgs ...map[string]interface{}) {
	fields := fieldsFromArgs(fieldArgs...)
	defer notifyErrorMonitoringProvider(err, fields)
	logError(sharedLogger.Panic(), err, fields)
}

// Fatal logs fatal error.
func Fatal(err error, fieldArgs ...map[string]interface{}) {
	fields := fieldsFromArgs(fieldArgs...)
	defer notifyErrorMonitoringProvider(err, fields)
	logError(sharedLogger.Fatal(), err, fields)
}

// Error logs an error.
func Error(err error, fieldArgs ...map[string]interface{}) {
	fields := fieldsFromArgs(fieldArgs...)
	defer notifyErrorMonitoringProvider(err, fields)
	logError(sharedLogger.Error(), err, fields)
}

// Warn logs a warning.
func Warn(err error, fieldArgs ...map[string]interface{}) {
	fields := fieldsFromArgs(fieldArgs...)
	logError(sharedLogger.Warn(), err, fields)
}

// Debug logs a debug message.
func Debug(msg string, fieldArgs ...map[string]interface{}) {
	fields := fieldsFromArgs(fieldArgs...)
	logMsg(sharedLogger.Debug(), msg, fields)
}

// Info logs an message.
func Info(msg string, fieldArgs ...map[string]interface{}) {
	fields := fieldsFromArgs(fieldArgs...)
	logMsg(sharedLogger.Info(), msg, fields)
}

// WithRequestID is a short-cut method for adding request ID to log.
func WithRequestID(requestID string) map[string]interface{} {
	return WithField(FieldNameRequestID, requestID)
}

// WithField is a short-cut method for adding a field to log.
func WithField(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: value,
	}
}

// private

func logMsg(event *zerolog.Event, msg string, fields map[string]interface{}) {
	event.Fields(fields).Msg(msg)
}

func logError(event *zerolog.Event, err error, fields map[string]interface{}) {
	if err == nil {
		log.Panic().Msg("nil err passed to logger")
	}

	event.Fields(fields).Caller().Stack().Err(err).Msg(err.Error())
}

func fieldsFromArgs(fields ...map[string]interface{}) map[string]interface{} {
	if fields == nil {
		return nil
	}

	if len(fields) > 1 {
		panic("unexpected number of fields")
	}

	return fields[0]
}

func notifyErrorMonitoringProvider(err error, fields map[string]interface{}) {
	if rollbar.Token() == "" {
		return
	}

	delete(fields, FieldNameRequestHeaders)
	delete(fields, FieldNameRequestBody)

	if _, okErr := err.(*errors.Err); okErr {
		if fields == nil {
			fields = make(map[string]interface{})
		}

		fields[FieldNameOriginalTrace] = errors.ErrorStack(err)
	}

	rollbar.ErrorWithStackSkipWithExtras(rollbar.ERR, err, zerolog.CallerSkipFrameCount, fields)
}
