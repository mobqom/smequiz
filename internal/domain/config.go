package domain

type AppConfig struct {
	host string `env:"PORT"`
	port string `env:"PORT"`
}
