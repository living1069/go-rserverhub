package app

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "log"
    "os"
    "rserverhub/models"
)

var DB *gorm.DB

func CreateConnection() {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    password := os.Getenv("DB_PASSWORD")
    user := os.Getenv("DB_USER")
    db := os.Getenv("DB_NAME")

    var uri = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, db)

    var err error
    DB, err = gorm.Open("postgres", uri)

    if err != nil {
        log.Fatal(err.Error())
    }

    if _, b := os.LookupEnv("DB_LOG"); b {
        DB.LogMode(true)
    }

    DB.AutoMigrate(&models.Configuration{})
}
