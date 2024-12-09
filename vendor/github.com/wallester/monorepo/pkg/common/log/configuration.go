package log

import (
	"github.com/juju/errors"
	"github.com/kelseyhightower/envconfig"
	loglevel "github.com/wallester/monorepo/pkg/common/log/level"
)

type Configuration struct {
	Environment  string `required:"true" envconfig:"global_environment"`
	RollbarToken string `envconfig:"rollbar_token"`

	Level       *loglevel.Level `envconfig:"log_level"`
	LogToFile   bool            `required:"true" envconfig:"log_to_file"`
	LogFileName string          `required:"true" envconfig:"log_file_name"`
	LogPrettify bool            `required:"true" envconfig:"log_prettify"` // whether it is needed to add indent to JSON log or not

	MaskSensitiveFields    bool `required:"true" envconfig:"mask_sensitive_fields"` // whether to mask sensitive fields like "password"
	XMLMaskSensitiveFields bool `envconfig:"xml_mask_sensitive_fields"`             // whether to mask sensitive fields like "password" in XML

	ConsoleWriter bool `envconfig:"console_writer"` // use a colored output instead of usual uncolored JSON output.
}

// NewConfiguration parses and returns a new Configuration.
func NewConfiguration(prefix string) (*Configuration, error) {
	var result Configuration
	if err := envconfig.Process(prefix, &result); err != nil {
		return nil, errors.Annotatef(err, "processing configuration with prefix failed: prefix=%s", prefix)
	}

	return &result, nil
}

//nolint:revive
type LogConfiguration = Configuration
