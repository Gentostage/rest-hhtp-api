package apiserver

//Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey string `toml:"session_key"`
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}

//../migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable&user=postgres&password=12345" up
//../migrate -path migrations -database "postgres://localhost/restapi_test?sslmode=disable&user=postgres&password=12345" up
