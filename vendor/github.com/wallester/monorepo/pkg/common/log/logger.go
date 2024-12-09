package log

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/juju/errors"
	"github.com/rollbar/rollbar-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	commonenv "github.com/wallester/monorepo/pkg/common/env"
	commonerrors "github.com/wallester/monorepo/pkg/common/errors"
	logapi "github.com/wallester/monorepo/pkg/common/log/api"
	"github.com/wallester/monorepo/pkg/common/log/hook"
	loglevel "github.com/wallester/monorepo/pkg/common/log/level"
	"github.com/wallester/monorepo/pkg/common/log/processevents"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger implements a logger.
type Logger struct {
	zeroLogger  zerolog.Logger
	childLogger bool
}

var (
	_ logapi.ILogger = (*Logger)(nil)
	_ io.Writer      = (*Logger)(nil)
)

func init() {
	zerolog.TimestampFieldName = FieldNameTime
	zerolog.LevelFieldName = FieldNameLevelString
	zerolog.MessageFieldName = FieldNameMsg
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.ErrorStackFieldName = FieldNameErrorStack
	zerolog.CallerSkipFrameCount = 4
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

// New initializes a new logger instance and returns a func to close it.
func New(name string, cfg Configuration) (*Logger, func()) {
	writer, closer := configureWriter(cfg)

	cfgLevel := loglevel.Info
	if cfg.Level != nil {
		cfgLevel = *cfg.Level
	}

	level := zerolog.Level(cfgLevel)

	zeroLogger := zerolog.New(writer).Level(level).With().Timestamp().Caller().Logger().
		Hook(hook.NewLevel()).
		Hook(hook.NewGitVersion()).
		Hook(hook.NewTracing()).
		Hook(hook.NewLoggerName(name))

	logger := &Logger{
		zeroLogger: zeroLogger,
	}

	ConfigureRollbar(cfg.RollbarToken, cfg.Environment)

	return logger, closer
}

func (l *Logger) ZeroLogger() zerolog.Logger {
	return l.zeroLogger
}

// NewChildLogger creates a new instance of current logger.
// It is used to create a new logger instance with additional fields.
// Example:
//
//	logger := log.NewChildLogger().String("key", "value")
//	logger.Info("message")
//	logger.Error("failed to do something")
//
// Output:
//
//	{"key":"value","level":"info","msg":"message","time":"2020-10-20T14:00:00+03:00"}
//	{"key":"value","level":"error","msg":"failed to do something","time":"2020-10-20T14:00:00+03:00"}
func (l *Logger) NewChildLogger() logapi.ILogger {
	return &Logger{
		zeroLogger:  l.zeroLogger.With().Logger(),
		childLogger: true,
	}
}

func (l *Logger) IsChildLogger() bool {
	return l.childLogger
}

func (l *Logger) Write(p []byte) (int, error) {
	l.zeroLogger.Error().Msg(string(p))
	return len(p), nil
}

func (l *Logger) Panic(err error, fieldArgs ...map[string]any) {
	if err == nil {
		return
	}

	fields := mergeCustomDataAndArgsFields(err, fieldArgs...)
	if rollbar.Token() != "" {
		defer notifyErrorMonitoringProvider(context.TODO(), rollbar.CRIT, err, fields)
	}

	logError(l.zeroLogger.Panic(), err, fields)
}

func (l *Logger) Fatal(err error, fieldArgs ...map[string]any) {
	if err == nil {
		return
	}

	fields := mergeCustomDataAndArgsFields(err, fieldArgs...)
	if rollbar.Token() != "" {
		defer notifyErrorMonitoringProvider(context.TODO(), rollbar.CRIT, err, fields)
	}

	logError(l.zeroLogger.Fatal(), err, fields)
}

// Error logs an error.
func (l *Logger) Error(err error, fieldArgs ...map[string]any) {
	if err == nil {
		return
	}

	fields := mergeCustomDataAndArgsFields(err, fieldArgs...)
	if rollbar.Token() != "" {
		defer notifyErrorMonitoringProvider(context.TODO(), rollbar.ERR, err, fields)
	}

	logError(l.zeroLogger.Error(), err, fields)
}

// Warn logs a warning.
func (l *Logger) Warn(err error, fieldArgs ...map[string]any) {
	if err != nil {
		fields := mergeCustomDataAndArgsFields(err, fieldArgs...)
		logError(l.zeroLogger.Warn(), err, fields)
	}
}

// Debug logs a debug message.
func (l *Logger) Debug(msg string, fieldArgs ...map[string]any) {
	fields := fieldsFromArgs(fieldArgs...)
	logMsg(l.zeroLogger.Debug(), msg, fields)
}

// Trace logs a trace message.
func (l *Logger) Trace(msg string, fieldArgs ...map[string]any) {
	fields := fieldsFromArgs(fieldArgs...)
	logMsg(l.zeroLogger.Trace(), msg, fields)
}

// Info logs an message.
func (l *Logger) Info(msg string, fieldArgs ...map[string]any) {
	fields := fieldsFromArgs(fieldArgs...)
	logMsg(l.zeroLogger.Info(), msg, fields)
}

func (l *Logger) NewEntry() logapi.IEntry {
	return &Entry{
		logger: l,
	}
}

var ErrMethodNotAllowedOnRootLogger = errors.New("method not allowed on root logger")

func (l *Logger) Bytes(name string, value []byte) logapi.ILogger {
	return l.updateChildLogger(func(c zerolog.Context) zerolog.Context {
		return c.Bytes(name, value)
	})
}

func (l *Logger) Int(name string, value int) logapi.ILogger {
	return l.updateChildLogger(func(c zerolog.Context) zerolog.Context {
		return c.Int(name, value)
	})
}

func (l *Logger) Int64(name string, value int64) logapi.ILogger {
	return l.updateChildLogger(func(c zerolog.Context) zerolog.Context {
		return c.Int64(name, value)
	})
}

func (l *Logger) String(name, value string) logapi.ILogger {
	return l.updateChildLogger(func(c zerolog.Context) zerolog.Context {
		return c.Str(name, value)
	})
}

func (l *Logger) Any(name string, value any) logapi.ILogger {
	return l.updateChildLogger(func(c zerolog.Context) zerolog.Context {
		return c.Any(name, value)
	})
}

func (l *Logger) UUID4(name string, value strfmt.UUID4) logapi.ILogger {
	return l.String(name, value.String())
}

func (l *Logger) Strings(name string, value []string) logapi.ILogger {
	return l.updateChildLogger(func(c zerolog.Context) zerolog.Context {
		return c.Strs(name, value)
	})
}

func (l *Logger) Field(name string, value any) logapi.ILogger {
	return l.Map(map[string]any{name: value})
}

func (l *Logger) Map(fields map[string]any) logapi.ILogger {
	return l.updateChildLogger(func(c zerolog.Context) zerolog.Context {
		return c.Fields(fields)
	})
}

func (l *Logger) Bool(name string, value bool) logapi.ILogger {
	return l.updateChildLogger(func(c zerolog.Context) zerolog.Context {
		return c.Bool(name, value)
	})
}

func WithRequestID(requestID string) map[string]any {
	return WithField(FieldNameRequestID, requestID)
}

func WithField(key string, value any) map[string]any {
	return map[string]any{
		key: value,
	}
}

// private

func (l *Logger) updateChildLogger(f func(c zerolog.Context) zerolog.Context) logapi.ILogger {
	if !l.childLogger {
		l.Panic(ErrMethodNotAllowedOnRootLogger)
	}

	l.zeroLogger.UpdateContext(f)
	return l
}

func logMsg(event *zerolog.Event, msg string, fields map[string]any) {
	event.Fields(fields).Msg(msg)
}

func logError(event *zerolog.Event, err error, fields map[string]any) {
	if err == nil {
		log.Panic().Msg("nil err passed to logger")
	}

	event.Fields(fields).Stack().Err(err).Send()
}

func mergeCustomDataAndArgsFields(err error, fieldArgs ...map[string]any) map[string]any {
	fields := fieldsFromCustomData(err, fieldsFromArgs(fieldArgs...))

	return fields
}

func fieldsFromCustomData(err error, fields map[string]any) map[string]any {
	if e, ok := errors.Cause(err).(commonerrors.IErrorWithCustomData); ok {
		fields = e.CustomData().Map(fields).Fields()
		return fieldsFromCustomData(e.Source(), fields)
	}

	return fields
}

func fieldsFromArgs(fields ...map[string]any) map[string]any {
	if fields == nil {
		return nil
	}

	if len(fields) > 1 {
		panic("unexpected number of fields")
	}

	return fields[0]
}

func configureWriter(cfg Configuration) (io.Writer, func()) {
	var writer io.Writer = os.Stdout
	if cfg.LogPrettify {
		writer = Writer{
			Out:    writer,
			Indent: true,
		}
	}

	if cfg.ConsoleWriter {
		writer = zerolog.NewConsoleWriter()
	}

	var closers []func()

	if cfg.LogToFile && cfg.LogFileName != "" {
		logFile, err := openLogFile(cfg)
		if err != nil {
			log.Fatal().Err(err).Msg("opening log file failed")
		}

		closers = append(closers, func() {
			if err := logFile.Close(); err != nil {
				log.Warn().Err(err).Msg("closing log file failed")
			}
		})

		writer = io.MultiWriter(logFile, writer)
	}

	if cfg.MaskSensitiveFields {
		writer = Writer{
			Out:          writer,
			ProcessEvent: processevents.MaskEvent,
		}
	}

	if cfg.XMLMaskSensitiveFields {
		writer = Writer{
			Out:          writer,
			ProcessEvent: processevents.XMLMaskEvent,
		}
	}

	if rollbar.Token() != "" {
		closers = append(closers, rollbar.Close)
	}

	closer := func() {
		for _, closer := range closers {
			closer()
		}
	}

	return writer, closer
}

func openLogFile(cfg Configuration) (io.WriteCloser, error) {
	if commonenv.IsDevelopment(cfg.Environment) {
		// Configure Lumberjack for log rotation
		lumberjackLogger := &lumberjack.Logger{
			Filename:   cfg.LogFileName,
			MaxSize:    10,    // Megabytes
			MaxBackups: 3,     // Number of old log files to keep
			MaxAge:     28,    // Days
			Compress:   false, // Compress old log files
		}

		return lumberjackLogger, nil
	}

	directory := filepath.Dir(cfg.LogFileName)
	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		return nil, errors.Annotatef(err, "creating directory failed: directory=%s", directory)
	}

	logFile, err := os.OpenFile(filepath.Clean(cfg.LogFileName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o600)
	if err != nil {
		return nil, errors.Annotatef(err, "opening log file failed: filePath=%s", cfg.LogFileName)
	}

	return logFile, nil
}
