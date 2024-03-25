package config

type EKSOptions struct {
	ClusterName string `yaml:"cluster_name" json:"cluster_name"`
	Region      string `yaml:"region" json:"region"`
	AccessKey   string `yaml:"access_key" json:"access_key"`
	SecretKey   string `yaml:"secret_key" json:"secret_key"`
}
