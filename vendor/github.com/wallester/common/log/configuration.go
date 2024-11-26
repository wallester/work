package log

import (
	"github.com/juju/errors"
	"github.com/kelseyhightower/envconfig"
	"github.com/wallester/common/env"
)

//nolint:maligned,golint
type Configuration struct {
	Environment  string `required:"true" envconfig:"global_environment"`
	Level        Level
	RollbarToken string `envconfig:"rollbar_token"`

	LogToFile   bool   `required:"true" envconfig:"log_to_file"`
	LogFileName string `required:"true" envconfig:"log_file_name"`
	LogPrettify bool   `required:"true" envconfig:"log_prettify"` // whether it is needed to add indent to JSON log or not

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

	if env.IsProduction(result.Environment) {
		result.Level = InfoLevel
	} else {
		result.Level = DebugLevel
	}

	return &result, nil
}

//nolint:golint
type LogConfiguration = Configuration
