package configs

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	JWTSecret      string
	PostgresConfig *PostgresConfig
}

type PostgresConfig struct {
	User     string
	Password string
	Database string
	Port     int
}

func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	port, err := strconv.Atoi(os.Getenv("POSTGRESQL_PORT"))
	if err != nil {
		panic("failed to load postgres port, err: " + err.Error())
	}
	return &Config{
		PostgresConfig: &PostgresConfig{
			User:     os.Getenv("POSTGRESQL_USERNAME"),
			Password: os.Getenv("POSTGRESQL_PASSWORD"),
			Database: os.Getenv("POSTGRESQL_DATABASE"),
			Port:     port,
		},
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
