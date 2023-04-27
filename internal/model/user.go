package model

import "time"

type User struct {
    ID          uint64 `gorm:"primary_key"`
    Name        string `gorm:"not null"`
    Email       string `gorm:"unique;not null"`
    KeyID       uint64 `gorm:"not null"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
