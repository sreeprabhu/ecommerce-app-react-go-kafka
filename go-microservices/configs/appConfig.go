package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort   string
	DBConnection string
	AppSecret    string
}

func SetupEnv() (cnf AppConfig, err error) {
	appEnv := os.Getenv("APP_ENV")
	appEnv = strings.TrimSpace(appEnv)

	fmt.Println("APP_ENV -", appEnv)
	/* 	you can simply run - $env:APP_ENV="dev"; go run main.go
	instead we have added a makeFile for local development
	*/
	if appEnv == "dev" {
		fmt.Println("App is running in development environment")
		// go get github.com/joho/godotenv
		// loads the .env file
		godotenv.Load()
		fmt.Println(os.Getenv("APP_ENV"))
	} else {
		fmt.Println("App is not runnning in development environment")
	}

	httpPort := os.Getenv("HTTP_PORT")

	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variable HTTP_PORT not found")
	}

	DBConnection := os.Getenv("DB_CONNECTION")
	if len(DBConnection) < 1 {
		return AppConfig{}, errors.New("env variable DB_CONNECTION not found")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("env variable APP_SECRET not found")
	}

	return AppConfig{ServerPort: httpPort, DBConnection: DBConnection, AppSecret: appSecret}, nil
}
