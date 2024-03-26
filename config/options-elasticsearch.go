package config

type ElasticsearchOptions struct {
	Active   bool   `yaml:"active" json:"active"`
	Host     string `yaml:"host" json:"host"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Index    string `yaml:"index" json:"index"`
}
