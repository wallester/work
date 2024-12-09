package devenv

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/juju/errors"
	commonerrors "github.com/wallester/monorepo/pkg/common/errors"
	"gopkg.in/yaml.v2"
)

// ServicesConfiguration represents the services configuration,
// as it is defined in the services.yml file.
// Basically it is a list of services, with their names and directories and ports.
type ServicesConfiguration struct {
	Services []*Service `yaml:"services"`
}

// NewServicesConfiguration reads the services.yml file and returns the ServicesConfiguration object.
func NewServicesConfiguration() (*ServicesConfiguration, error) {
	filename := filepath.Join(MonorepoPath(), "services.yml")

	data, err := os.ReadFile(filename) //#nosec:G304
	if err != nil {
		return nil, errors.Annotate(err, "failed to read services file")
	}

	var cfg ServicesConfiguration
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, errors.Annotate(err, "failed to parse services file")
	}

	// Sort services by directory name
	sort.Slice(cfg.Services, func(i, j int) bool {
		return cfg.Services[i].Directory < cfg.Services[j].Directory
	})

	if err := cfg.Validate(); err != nil {
		return nil, errors.Annotate(err, "failed to validate services file")
	}

	return &cfg, nil
}

// Validate checks if the services configuration is valid, e.g. if the service names and directories are unique.
func (cfg *ServicesConfiguration) Validate() error {
	return commonerrors.Last(ServicesToMap(cfg.Services))
}
