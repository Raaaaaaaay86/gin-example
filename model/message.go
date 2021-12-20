package model

type Message struct {
	Model
	Content   string
	UserID    uint
	ChannelID uint
	User      User
	Channel   Channel
}
