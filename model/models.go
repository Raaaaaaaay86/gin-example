package model

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func Setup() {
	var err error
	db, err = gorm.Open(mysql.Open("root:eee333@tcp(127.0.0.1:3307)/playground?charset=utf8&parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
}

func GetDB() *gorm.DB {
	return db
}
