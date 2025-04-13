package configuration

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Application struct {
		Mode      string
		SecretKey []byte
	}
	Server struct {
		Port string
	}
	Database struct {
		Host        string
		Port        string
		Username    string
		Password    string
		Database    string
		DevDatabase string
	}
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Err loading .env file")
	}
	_, err = os.Stat(".env.local")
	if err == nil {
		err := godotenv.Overload(".env.local")
		if err != nil {
			panic("Err loading .env.local file")
		}
	}
	AppSecretKey := []byte(os.Getenv("APP_SECRET_KEY"))
	if len(AppSecretKey) == 0 {
		panic("APP_SECRET_KEY is not set")
	}
	return &Config{
		Application: struct {
			Mode      string
			SecretKey []byte
		}{
			Mode:      os.Getenv("APP_MODE"),
			SecretKey: AppSecretKey,
		},
		Server: struct{ Port string }{
			Port: os.Getenv("API_PORT"),
		},
		Database: struct {
			Host        string
			Port        string
			Username    string
			Password    string
			Database    string
			DevDatabase string
		}{
			Host:        os.Getenv("DB_HOST"),
			Port:        os.Getenv("DB_PORT"),
			Username:    os.Getenv("DB_USER"),
			Password:    os.Getenv("DB_PASSWORD"),
			Database:    os.Getenv("DB_NAME"),
			DevDatabase: os.Getenv("DB_DEV_NAME"),
		},
	}
}
