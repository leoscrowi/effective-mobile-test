package config

import "os"

type AppConfig struct {
	Port string `default:"8080"`
}

type Config struct {
	DatabaseConfig DatabaseConfig
	AppConfig      AppConfig
}

type DatabaseConfig struct {
	Host     string `default:"localhost"`
	Port     string `default:"5432"`
	User     string `default:"postgres"`
	Password string `default:"postgres"`
	Name     string `default:"subscription_db"`
	SslMode  string `default:"disable"`
}

func MustLoad() *Config {
	return &Config{
		DatabaseConfig: DatabaseConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Name:     os.Getenv("POSTGRES_DB"),
			SslMode:  os.Getenv("POSTGRES_SSL_MODE"),
		},
		AppConfig: AppConfig{
			Port: os.Getenv("APP_PORT"),
		},
	}
}
