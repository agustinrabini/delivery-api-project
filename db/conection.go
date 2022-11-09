package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DbConect() *sql.DB {

	conString := getEnvVars()
	dbProducts, err := sql.Open("mysql", conString)

	if err != nil {
		log.Fatal("Unnable to connect DB")
	}

	return dbProducts
}

func getEnvVars() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("USR")
	password := os.Getenv("PASSWD")
	hostname := os.Getenv("PORT")
	dbName := os.Getenv("SCHEMA")

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
