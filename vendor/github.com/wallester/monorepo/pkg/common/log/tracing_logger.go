package log

import (
	"fmt"

	"github.com/juju/errors"
)

type TracingLogger struct{}

func (*TracingLogger) Debugf(format string, args ...any) {
	Debug(fmt.Sprintf(format, args...))
}

func (*TracingLogger) Errorf(format string, args ...any) {
	Warn(errors.Errorf(format, args...))
}
