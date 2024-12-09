package log

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/juju/errors"
	"github.com/rollbar/rollbar-go"
	"github.com/rs/zerolog"
	"github.com/wallester/monorepo/pkg/common/git"
)

func ConfigureRollbar(rollbarToken, environment string) {
	if rollbarToken == "" {
		return
	}

	rollbar.SetToken(rollbarToken)
	rollbar.SetEnvironment(environment)
	rollbar.SetCodeVersion(git.Version())
	rollbar.SetTelemetry(
		rollbar.EnableLoggerTelemetry(),
		rollbar.EnableNetworkTelemetry(&http.Client{
			Timeout: time.Second * 30,
		}),
		rollbar.EnableNetworkTelemetryRequestHeaders(),
		rollbar.EnableNetworkTelemetryResponseHeaders(),
	)
	rollbar.SetCaptureIp(rollbar.CaptureIpFull)
}

// private

func notifyErrorMonitoringProvider(ctx context.Context, level string, err error, fields map[string]any) {
	delete(fields, FieldNameRequestHeaders)
	delete(fields, FieldNameRequestBody)

	if _, okErr := err.(*errors.Err); okErr {
		if fields == nil {
			fields = make(map[string]any)
		}

		fields[FieldNameOriginalTrace] = errors.ErrorStack(err)
	}

	rollbar.SetTransform(func(data map[string]any) {
		if requestID, ok := fields[FieldNameRequestID].(string); ok && strfmt.IsUUID4(requestID) {
			data["uuid"] = requestID
		}
	})

	rollbar.ErrorWithStackSkipWithExtrasAndContext(ctx, level, err, zerolog.CallerSkipFrameCount, fields)
}
