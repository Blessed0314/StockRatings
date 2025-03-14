package db

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    password := os.Getenv("DB_PASSWORD")
    if password != "" {
        password = " password=" + password
    }
    
    dsn := "host=" + os.Getenv("DB_HOST") +
        " user=" + os.Getenv("DB_USER") +
        password +
        " dbname=" + os.Getenv("DB_NAME") +
        " port=" + os.Getenv("DB_PORT") +
        " sslmode=disable"

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database", err)
    }

    log.Println("Conexión exitosa a CockroachDB 🐓!")

}