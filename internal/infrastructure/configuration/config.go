package configuration

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Application struct {
		Mode string
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

var AppSecretKey []byte

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	AppSecretKey = []byte(os.Getenv("APP_SECRET_KEY"))
	if len(AppSecretKey) == 0 {
		panic("APP_SECRET_KEY is not set")
	}
	return &Config{
		Application: struct{ Mode string }{
			Mode: os.Getenv("APP_MODE"),
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

func (c *Config) GetPostgresDsn() string {
	dsn := "postgres://" + c.Database.Username + ":" + c.Database.Password + "@"
	dsn += c.Database.Host + ":" + c.Database.Port + "/" + c.Database.Database + "?sslmode=disable"
	return dsn
}
