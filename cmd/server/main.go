package main

import (
    "key-app-go/internal/db"
)

func main() {
    err := db.AutoMigrate()
    if err != nil {
        panic(err)
    }
}