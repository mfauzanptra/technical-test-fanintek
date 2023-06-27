package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var JWT_KEY string = ""

type AppConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
	jwtKey string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}

	err := godotenv.Load("config.env")
	if err != nil {
		fmt.Println("Error saat baca env", err.Error())
		return nil
	}

	app.DBUser = os.Getenv("DBUSER")
	app.DBPass = os.Getenv("DBPASS")
	app.DBHost = os.Getenv("DBHOST")
	readData := os.Getenv("DBPORT")
	app.DBPort, err = strconv.Atoi(readData)
	if err != nil {
		fmt.Println("Error saat convert", err.Error())
		return nil
	}
	app.DBName = os.Getenv("DBNAME")
	app.jwtKey = os.Getenv("JWTKEY")

	JWT_KEY = app.jwtKey

	return &app
}
