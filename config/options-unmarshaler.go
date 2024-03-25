package config

import (
	"encoding/json"
	"errors"

	"gopkg.in/yaml.v3"
)

func loadOptions(content []byte, out any) error {
	err := yaml.Unmarshal(content, out)
	if err == nil {
		return nil
	}

	err = json.Unmarshal(content, out)
	if err == nil {
		return nil
	}

	return errors.New("unable to unmarshal configuration file")
}
