package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(ac AppConfig) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", ac.DBHost, ac.DBPort, ac.DBUser, ac.DBPass, ac.DBName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}
