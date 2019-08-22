package store

// Config ...
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
