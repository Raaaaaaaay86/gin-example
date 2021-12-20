package user_service

import (
	"example/gorm-practice/model"
	"fmt"
)

func GetAllUser() []*model.User {
	return model.GetAllUser()
}

func CreateUser(username string, email string) *model.User {
	return model.CreateUser(username, email)
}

func FindUserById(id uint) (*model.User, error) {
	user, err := model.FindUserById(id)
	fmt.Println(user)
	return user, err
}
