package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func MyPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		os.Setenv(port, "8080")
		port = "8080"
	}
	return ":" + port, nil
}

func Connectdb() (*sql.DB, error) {
	path := os.Getenv("DB_URL") + os.Getenv("DB_NAME")
	db, errdb := sql.Open(os.Getenv("DB_TYPE"), path)
	if errdb != nil {
		return nil, errdb
	}
	err := db.Ping()
	return db, err
}
