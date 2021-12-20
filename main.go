package main

import (
	"example/gorm-practice/model"
	"example/gorm-practice/router"
)

func init() {
	model.Setup()
}

func main() {
	router := router.InitRouter()
	router.Run()
}
