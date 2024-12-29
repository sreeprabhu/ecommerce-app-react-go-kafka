package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (cnf AppConfig, err error) {

	/* 	you can simply run - $env:APP_ENV="dev"; go run main.go
	instead we have added a makeFile for local development
	*/
	if os.Getenv("APP_ENV") == "dev" {
		// go get github.com/joho/godotenv
		// loads the .env file
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variable HTTP_PORT not found")
	}

	return AppConfig{ServerPort: httpPort}, nil
}
