package router

import (
	"example/gorm-practice/router/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("/user/all", api.GetAllUser)
	router.GET("/user/:id", api.FindUserById)
	router.POST("/user/add", api.CreateUser)
	router.GET("/test", api.Test)
	return router
}
