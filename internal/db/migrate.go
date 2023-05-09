package db

import (

    "fmt"
    "github.com/jinzhu/gorm"
    "key-app-go/internal/model"
        _ "github.com/go-sql-driver/mysql"
)

func AutoMigrate() error {
    db, err := gorm.Open("mysql", "keyuser:keypassword@tcp(172.23.0.2:3305)/db?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        return fmt.Errorf("error al conectar a la BD: %v", err)
    }

    if err := db.AutoMigrate(&model.Key{}, &model.User{}).Error; err != nil {
        return fmt.Errorf("error al migrar los modelos: %v", err)
    }


    fmt.Println("Migraciones realizada con éxito")
    return nil
}


