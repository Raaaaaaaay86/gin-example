package model

import (
	"fmt"
)

type User struct {
	Model
	Email    string `json:"email"`
	Username string `json:"username"`
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
