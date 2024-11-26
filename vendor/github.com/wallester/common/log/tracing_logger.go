package log

import (
	"fmt"

	"github.com/juju/errors"
)

type TracingLogger struct{}

func (*TracingLogger) Debugf(format string, args ...interface{}) {
	Debug(fmt.Sprintf(format, args...))
}

func (*TracingLogger) Errorf(format string, args ...interface{}) {
	Warn(errors.New(fmt.Sprintf(format, args...)))
}
