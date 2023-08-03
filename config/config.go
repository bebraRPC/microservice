package config

type Config struct {
	Redis    RedisConfig
	RabbitMQ RabbitMQ
	Logger   Logger
	MongoDB  MongoDB
}

type RabbitMQ struct {
	Host           string
	Port           string
	User           string
	Password       string
	Exchange       string
	Queue          string
	RoutingKey     string
	ConsumerTag    string
	WorkerPoolSize int
}

type MongoDB struct {
	URI      string
	User     string
	Password string
	DB       string
}

type Logger struct {
	Level string
}

type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultDB string
	MinIdleConn    int
	PoolSize       int
	PoolTimeout    int
	Password       string
	DB             int
}
