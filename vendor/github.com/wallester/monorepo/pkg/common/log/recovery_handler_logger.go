package log

import (
	"runtime/debug"

	"github.com/juju/errors"
	"github.com/rollbar/rollbar-go"
	"github.com/rs/zerolog"
)

func RecoveredPanic(r any) {
	err := errors.Errorf("%v", r)

	if rollbar.Token() != "" {
		fields := make(map[string]any)
		fields[FieldNameOriginalTrace] = string(debug.Stack())

		rollbar.ErrorWithStackSkipWithExtras(rollbar.CRIT, err, zerolog.CallerSkipFrameCount, fields)
		rollbar.Wait()
	} else {
		Error(err)
	}
}
