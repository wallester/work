package log

import (
	commonstrings "github.com/wallester/monorepo/pkg/common/strings"
)

func TrimLength(s string) string { return commonstrings.TrimLength(s, 1024) }
