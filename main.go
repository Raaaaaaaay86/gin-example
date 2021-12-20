package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	Name        string
	Description string
}

type User struct {
	gorm.Model
	Email    string
	Username string
}

type Message struct {
	gorm.Model
	Content   string
	UserID    uint
	ChannelID uint
	User      User
	Channel   Channel
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&Channel{}, &User{}, &Message{})
	// seed(db)
}

func seed(db *gorm.DB) {
	channels := []Channel{
		{Name: "General", Description: "General Discussions"},
		{Name: "Off-Topic", Description: "Weird stuff goes here"},
		{Name: "Suggestions", Description: "Video suggestions go here"},
	}
	for _, channel := range channels {
		db.Create(&channel)
	}
	users := []User{
		{Email: "test@test.com", Username: "Joe420"},
		{Email: "yes@yes.com", Username: "Bob"},
	}
	for _, user := range users {
		db.Create(&user)
	}
	var generalChat, suggestionsChat Channel
	db.First(&generalChat, "Name = ?", "General")
	db.First(&suggestionsChat, "Name = ?", "Suggestions")
	var joe, bob User
	db.First(&joe, "Username = ?", "Joe420")
	db.First(&bob, "Username = ?", "bob")

	messages := []Message{
		{Content: "Hello!", Channel: generalChat, User: joe},
		{Content: "What up", Channel: generalChat, User: bob},
		{Content: "Make more go videos", Channel: suggestionsChat, User: joe},
	}
	for _, message := range messages {
		db.Create(&message)
	}
}

func main() {
	db, err := gorm.Open(mysql.Open("root:eee333@tcp(127.0.0.1:3307)/playground?charset=utf8&parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}

	setup(db)

	var users []User
	db.Find(&users)
	for _, user := range users {
		fmt.Println("Email: ", user.Email, " Username: ", user.Username)
	}

	var messages []Message
	db.Find(&messages)
	for _, message := range messages {
		fmt.Println("Message: ", message.Content, "Sender: ", message.UserID)
	}

	doError(db)
}

func doError(db *gorm.DB) {
	var fred User

	if err := db.Where("username = ?", "Fred").First(&fred).Error; err != nil {
		log.Fatalf("Error when loading User: %s", err)
	}

	// result := db.Where("username = ?", "Fred").First(&fred)
	// if result.Error != nil {
	// 	log.Fatalf("Error when loading User: %s", result.Error)
	// }
}
