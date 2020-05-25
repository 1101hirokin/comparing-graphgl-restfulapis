package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Protocol string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "hiroki",
			Password: "hiroki_pass",
			Name:     "furukawa_seminer",
			Protocol: "tcp(db:3306)",
			Charset:  "utf8",
		},
	}
}
