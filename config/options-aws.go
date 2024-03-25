package config

import (
	"fmt"
	"strings"

	"github.com/papaya147/kube-logger/util"
)

type EKSOptions struct {
	ClusterName string `yaml:"cluster_name" json:"cluster_name"`
	Region      string `yaml:"region" json:"region"`
	AccessKey   string `yaml:"access_key" json:"access_key"`
	SecretKey   string `yaml:"secret_key" json:"secret_key"`
}

func NewEKSOptions(paths ...string) (*EKSOptions, error) {
	for _, path := range paths {
		content, err := util.GetFileContents(path)
		if err != nil {
			continue
		}

		var options EKSOptions
		if err := loadOptions(content, &options); err != nil {
			continue
		}

		return &options, nil
	}

	return nil, fmt.Errorf("files %s could not be loaded as options", strings.Join(paths, ", "))
}
