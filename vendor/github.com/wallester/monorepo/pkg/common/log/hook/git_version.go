package hook

import (
	"github.com/rs/zerolog"
	"github.com/wallester/monorepo/pkg/common/git"
)

type GitVersion struct{}

func NewGitVersion() GitVersion {
	return GitVersion{}
}

func (h GitVersion) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	// Add code version (git commit) to the logger
	e.Str(gitVersionEvent, git.Version())
}

// private

// gitVersionEvent is a constant which represents string field git-version in zerolog event.
const gitVersionEvent = "git-version"
