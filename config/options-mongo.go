package config

type MongoOptions struct {
	Active        bool   `yaml:"active" json:"active"`
	ConnectionURI string `yaml:"connection_uri" json:"connection_uri"`
	Database      string `yaml:"database" json:"database"`
	Collection    string `yaml:"collection" json:"collection"`
}
