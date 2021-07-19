package entities

import "gorm.io/gorm"

type ConsumersLineAccounts struct {
	gorm.Model
	CreateBy          string
	LineId            string
	UserId            string
	DisplayName       string
	PictureUrl        string
	ConsumerSegmentId int
}
