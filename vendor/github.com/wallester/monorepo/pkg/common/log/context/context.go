package context

import (
	"context"

	logapi "github.com/wallester/monorepo/pkg/common/log/api"
	loglevel "github.com/wallester/monorepo/pkg/common/log/level"
)

func WithLogger(ctx context.Context, logger logapi.ILogger) context.Context {
	return context.WithValue(ctx, contextKeyLogger, logger)
}

func Logger(ctx context.Context) logapi.ILogger {
	res, ok := ctx.Value(contextKeyLogger).(logapi.ILogger)
	if ok {
		return res
	}

	return nil
}

func WithQueryLogLevel(ctx context.Context, level loglevel.Level) context.Context {
	return context.WithValue(ctx, contextKeyQueryLogLevel, level)
}

func WithQueryLogLevelDebug(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKeyQueryLogLevel, loglevel.Debug)
}

func QueryLogLevel(ctx context.Context) (loglevel.Level, bool) {
	value, ok := ctx.Value(contextKeyQueryLogLevel).(loglevel.Level)
	return value, ok
}

// private

type contextKey string

const (
	contextKeyQueryLogLevel contextKey = "QueryLogLevel"
	contextKeyLogger        contextKey = "Logger"
)
