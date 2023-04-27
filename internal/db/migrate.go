package db

import (
    "fmt"
    "log"
    "gorm.io/gorm"
    "key-app-go/internal/model"
    "github.com/go-sql-driver/mysql"
     "github.com/jinzhu/gorm"
)

func AutoMigrate() error {

    dbHost := "172.17.0.1"
    dbPort := "3305"
    dbName := "keyd"
    dbUser := "keyuser"
    dbPassword := "keypassword"

    dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

    db, err := gorm.Open("mysql", dbURI)
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    defer db.Close()

    db.AutoMigrate(&model.Key{})
    db.AutoMigrate(&model.User{})

    log.Println("Migration has been successfully completed!")

}