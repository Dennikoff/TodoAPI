package apiserver

type Config struct {
	Bind_addr   string `toml:"bind_addr"`
	Loglevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}
