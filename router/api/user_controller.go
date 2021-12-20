package api

import (
	user_service "example/gorm-practice/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"data": user_service.GetAllUser(),
		},
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
		gin.H{
			"data": user_service.CreateUser(requestBody.Username, requestBody.Email),
		},
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
