package model

import "time"

type Key struct {
    ID          uint64 `gorm:"primary_key"`
    Name        string `gorm:"unique;not null"`
    Description string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}