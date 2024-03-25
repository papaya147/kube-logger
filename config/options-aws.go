package config

import (
	"github.com/papaya147/kube-logger/util"
)

type EKSOptions struct {
	ClusterName string `yaml:"cluster_name" json:"cluster_name"`
	Region      string `yaml:"region" json:"region"`
	AccessKey   string `yaml:"access_key" json:"access_key"`
	SecretKey   string `yaml:"secret_key" json:"secret_key"`
}

func NewEKSOptions(path string) (*EKSOptions, error) {
	content, err := util.GetFileContents(path)
	if err != nil {
		return nil, err
	}

	var options EKSOptions
	if err := loadOptions(content, &options); err != nil {
		return nil, err
	}

	return &options, nil
}
