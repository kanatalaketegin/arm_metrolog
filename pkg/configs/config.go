package configs

type Config struct {
	DB *DBConfig
}
type DBConfig struct {
	Host     string
	Port     int
	UserName string
	DBName   string
	Password string
	SslMode  string
}

func NewConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host:     "localhost",
			Port:     5431,
			UserName: "postgres",
			Password: "password",
			DBName:   "postgres",
			SslMode:  "default",
		},
	}
}
