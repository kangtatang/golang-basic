package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// ConnectDB menghubungkan ke database PostgreSQL
func ConnectDB() (*gorm.DB, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, fmt.Errorf("error loading .env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbName, dbPassword)

    db, err := gorm.Open("postgres", dbURI)
    if err != nil {
        return nil, err
    }

    return db, nil
}
