package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetConfig() string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error reading env file")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	return dsn
}
