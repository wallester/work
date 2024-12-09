package devenv

import (
	"os"
	"path/filepath"
)

func GoPath() string                  { return os.Getenv("GOPATH") }
func WallesterPath() string           { return filepath.Join(GoPath(), "src", "github.com", "wallester") }
func MonorepoPath() string            { return filepath.Join(WallesterPath(), "monorepo") }
func AutomationPath() string          { return filepath.Join(MonorepoPath(), "automation") }
func CoveragePath() string            { return filepath.Join(AutomationPath(), ".coverage") }
func IntegrationCoveragePath() string { return filepath.Join(CoveragePath(), "integration") }
func MergedCoveragePath() string      { return filepath.Join(CoveragePath(), "merged") }
