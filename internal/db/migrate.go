package db

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func AutoMigrate() error {
    dsn := "keyuser:keypassword@tcp(127.0.0.1:3305)/keyd?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    err = db.AutoMigrate(&model.Key{}, &model.User{})
    if err != nil {
        return err
    }

    return nil
}
