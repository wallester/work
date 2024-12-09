package work

import (
	"github.com/juju/errors"
	commonlog "github.com/wallester/monorepo/pkg/common/log"
)

func logError(key string, err error) {
	commonlog.Error(errors.Annotate(err, key))
}
