package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/papaya147/kube-logger/util"
	"gopkg.in/yaml.v3"
)

type Options struct {
	Namespaces      []string     `yaml:"namespaces" json:"namespaces"`
	ClusterProvider string       `yaml:"cluster_provider" json:"cluster_provider"`
	Writer          string       `yaml:"writer" json:"writer"`
	EKSOptions      EKSOptions   `yaml:"eks" json:"eks"`
	MongoOptions    MongoOptions `yaml:"mongo" json:"mongo"`
}

func NewOptions(paths ...string) (*Options, error) {
	for _, path := range paths {
		content, err := util.GetFileContents(path)
		if err != nil {
			continue
		}

		var options Options
		if err := loadOptions(content, &options); err != nil {
			continue
		}

		return &options, nil
	}

	return nil, fmt.Errorf("files %s could not be loaded as options", strings.Join(paths, ", "))
}

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
