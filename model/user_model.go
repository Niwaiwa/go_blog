package model

import (
	"errors"
	"go_blog/db"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Account   string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(username, account, passwrod string) (*User, error) {
	user := User{Username: username, Account: account, Password: passwrod}
	result := db.DBm.Create(&user)
	if err := result.Error; err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("user id: ", user.ID, "raw: ", result.RowsAffected)
	return &user, nil
}

func GetUser(account string) (*User, error) {
	user := User{Account: account}
	result := db.DBm.First(&user)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("user not found")
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	// if result.RowsAffected == 0 {
	// 	log.Println("account not found")
	// 	return nil, nil
	// }
	log.Println("user id: ", user.ID, "username: ", user.Username, "account: ", user.Account, "raw: ", result.RowsAffected)
	return &user, nil
}
