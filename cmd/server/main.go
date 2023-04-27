package main

import (
    "internal/db"

)

func main() {
    err := db.AutoMigrate()
    if err != nil {
        panic(err)
    }
}