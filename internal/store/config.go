package store

type Config struct {
	Url string `toml:"db_url"`
}

func NewConfig() *Config {
	return &Config{}
}
