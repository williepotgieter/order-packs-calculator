package config

type ApiConfig struct {
	Port uint16 `env:"PORT, default=8080"`
}

type AppConfig struct {
	Prod bool       `env:"PROD, default=false"`
	Api  *ApiConfig `env:", prefix=API_"`
}
