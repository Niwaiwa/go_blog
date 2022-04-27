package model

import (
	"go_blog/db"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Account   string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Create(username, account, passwrod string) {
	user := User{Username: username, Account: account, Password: passwrod}
	result := db.DBm.Create(&user)
}
