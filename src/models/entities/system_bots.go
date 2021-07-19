package entities

import "gorm.io/gorm"

type SystemBots struct {
	gorm.Model
	LineId string
	UserId string
	CreateBy string
	AgentId int
	ChannelSecret string
	ChannelAccessToken string
	WebhookUrl string
	DisplayName string
	PictureUrl string
}