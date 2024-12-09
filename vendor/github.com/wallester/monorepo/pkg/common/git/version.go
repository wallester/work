package git

import (
	"os"
	"runtime/debug"
	"sync"
)

var (
	version string
	once    sync.Once
)

// Version returns Git version. It is set from Go's build info when package is imported.
func Version() string {
	once.Do(func() {
		if version == "" {
			version = getBuildVersion()
		}
	})

	return version
}

// SetBuildInfoSource sets build info source func to use in tests for lazy-loading.
func SetBuildInfoSource(infoFunc func() (info *debug.BuildInfo, ok bool)) {
	buildInfoSource = infoFunc
}

// SetVersion sets Git version to v to use in tests.
func SetVersion(v string) {
	version = v
}

// private

var buildInfoSource = debug.ReadBuildInfo

func getBuildVersion() string {
	info, ok := buildInfoSource()
	if !ok {
		return ""
	}

	for _, setting := range info.Settings {
		if setting.Key == "vcs.revision" && setting.Value != "" {
			return setting.Value
		}
	}

	// fallback to use CODE_VERSION env variable instead.
	return os.Getenv("CODE_VERSION")
}
