package api

import (
	user_service "example/gorm-practice/service/user"
	"example/gorm-practice/utils/custom_response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		custom_response.Build(
			user_service.GetAllUser(),
			"ok",
		),
	)
}

func CreateUser(c *gin.Context) {
	type CreateUserRequest struct {
		Username string
		Email    string
	}
	var requestBody CreateUserRequest
	c.BindJSON(&requestBody)

	c.JSON(
		http.StatusOK,
		custom_response.Build(
			user_service.CreateUser(requestBody.Username, requestBody.Email),
			"ok",
		),
	)
}

func FindUserById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, _ := user_service.FindUserById(uint(id))

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": user,
		},
	)
}
