package init_db

import (
	"example/gorm-practice/model"

	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	db.AutoMigrate(&model.Channel{}, &model.User{}, &model.Message{})
	seed(db)
}

// Init Empty Database's data
func seed(db *gorm.DB) {
	channels := []model.Channel{
		{Name: "General", Description: "General Discussions"},
		{Name: "Off-Topic", Description: "Weird stuff goes here"},
		{Name: "Suggestions", Description: "Video suggestions go here"},
	}
	for _, channel := range channels {
		db.Create(&channel)
	}
	users := []model.User{
		{Email: "test@test.com", Username: "Joe420"},
		{Email: "yes@yes.com", Username: "Bob"},
	}
	for _, user := range users {
		db.Create(&user)
	}
	var generalChat, suggestionsChat model.Channel
	db.First(&generalChat, "Name = ?", "General")
	db.First(&suggestionsChat, "Name = ?", "Suggestions")
	var joe, bob model.User
	db.First(&joe, "Username = ?", "Joe420")
	db.First(&bob, "Username = ?", "bob")

	messages := []model.Message{
		{Content: "Hello!", Channel: generalChat, User: joe},
		{Content: "What up", Channel: generalChat, User: bob},
		{Content: "Make more go videos", Channel: suggestionsChat, User: joe},
	}
	for _, message := range messages {
		db.Create(&message)
	}
}
