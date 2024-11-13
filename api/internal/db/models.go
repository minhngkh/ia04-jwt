package db

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique;index"`
	Password  string
	CreatedAt time.Time
}
