package log

import (
	"context"
	"os"
	"sync"

	"github.com/juju/errors"
	"github.com/rs/zerolog"
	logapi "github.com/wallester/monorepo/pkg/common/log/api"
	logcontext "github.com/wallester/monorepo/pkg/common/log/context"
	"github.com/wallester/monorepo/pkg/common/log/hook"
)

func SharedLogger() logapi.ILogger {
	initializeSharedLoggerOnce.Do(func() {
		if deprecatedSharedLogger == nil {
			deprecatedSharedLogger = &Logger{
				zeroLogger: zerolog.New(os.Stdout).
					With().
					Timestamp().
					Caller().
					Logger().
					Hook(hook.NewLevel()).
					Hook(hook.NewGitVersion()).
					Hook(hook.NewTracing()).
					Hook(hook.NewLoggerName("shared-default")),
			}

			deprecatedSharedLogger.Warn(errors.New("logger not configured, using default configuration"))
		}
	})
	return deprecatedSharedLogger
}

func ConfigureSharedLogger(cfg Configuration) func() {
	configureSharedLoggerMutex.Lock()
	defer configureSharedLoggerMutex.Unlock()

	logger, closer := New("shared-configured", cfg)
	deprecatedSharedLogger = logger
	deprecatedSharedLogger.Debug("shared logger configured")
	return closer
}

func LoggerFromContext(ctx context.Context) logapi.ILogger {
	if ctx != nil {
		if logger := logcontext.Logger(ctx); logger != nil {
			return logger
		}
	}

	return SharedLogger()
}

func Panic(err error, fieldArgs ...map[string]any) {
	SharedLogger().Panic(err, fieldArgs...)
}

func Fatal(err error, fieldArgs ...map[string]any) {
	SharedLogger().Fatal(err, fieldArgs...)
}

func Error(err error, fieldArgs ...map[string]any) {
	SharedLogger().Error(err, fieldArgs...)
}

func Warn(err error, fieldArgs ...map[string]any) {
	SharedLogger().Warn(err, fieldArgs...)
}

func Debug(msg string, fieldArgs ...map[string]any) {
	SharedLogger().Debug(msg, fieldArgs...)
}

func Trace(msg string, fieldArgs ...map[string]any) {
	SharedLogger().Trace(msg, fieldArgs...)
}

func Info(msg string, fieldArgs ...map[string]any) {
	SharedLogger().Info(msg, fieldArgs...)
}

func NewEntry() *Entry {
	return &Entry{
		logger: SharedLogger(),
		fields: make(map[string]any),
	}
}

// private

var (
	deprecatedSharedLogger     logapi.ILogger
	initializeSharedLoggerOnce sync.Once
	configureSharedLoggerMutex sync.Mutex
)
