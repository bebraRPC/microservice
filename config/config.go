package config

type Config struct {
	MongoDB MongoDB
}

type MongoDB struct {
	URI      string
	User     string
	Password string
	DB       string
}
