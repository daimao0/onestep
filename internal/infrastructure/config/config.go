package config

type DatabaseConfig struct {
	Host string `yaml:"host"`

	Port string `yaml:"port"`

	User string `yaml:"user"`

	Password string `yaml:"password"`

	MaxOpenConn int `yaml:"maxOpenConn"`

	MaxIdleConn int `yaml:"maxIdleConn"`

	MaxLifetime int `yaml:"maxLifetime"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

var AppConfig = &Config{}
