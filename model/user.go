package model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Model
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.Username == "admin" {
		return errors.New("invalid username")
	}
	return
}

func GetAllUser() []*User {
	var users []*User
	if err := db.Find(&users).Error; err != nil {
		fmt.Println(err.Error())
	}
	return users
}

func CreateUser(username string, email string) *User {
	var user = User{Username: username, Email: email}

	if err := db.Create(&user).Error; err != nil {
		fmt.Println(err.Error())
	}

	return &user
}

func FindUserById(id uint) (*User, error) {
	var user *User
	if err := db.Find(&user, id).Error; err != nil {
		fmt.Println(err.Error())
	}
	return user, nil
}
