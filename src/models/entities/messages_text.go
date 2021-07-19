package entities

import "gorm.io/gorm"

type MessagesText struct {
	gorm.Model
	CreateBy string
	AgentId  int
	Title    string
	Content  string
	Status   int
}
