package log

import (
	"runtime/debug"

	"github.com/juju/errors"
	rollbar "github.com/rollbar/rollbar-go"
	"github.com/rs/zerolog"
)

func RecoveredPanic(r interface{}) {
	err := errors.Errorf("%v", r)

	if rollbar.Token() != "" {
		fields := make(map[string]interface{})
		fields[FieldNameOriginalTrace] = string(debug.Stack())

		rollbar.ErrorWithStackSkipWithExtras(rollbar.CRIT, err, zerolog.CallerSkipFrameCount, fields)
		rollbar.Wait()
	} else {
		Error(err)
	}
}
