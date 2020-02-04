package apiserver

//Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
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
// curl -d '{"email":"leontat20@gmail.com","password":"123456"}' -H "Content-Type: application/json" -X POST http://localhost:8080/users
// curl -c cookie.txt -d '{"email":"leontat20@gmail.com","password":"123456"}' -H "Content-Type: application/json" -X POST http://localhost:8080/sessions  --Запись Cookie
// curl -b cookie.txt http://localhost:8080/private/whoami  --Испольщование Cookie
