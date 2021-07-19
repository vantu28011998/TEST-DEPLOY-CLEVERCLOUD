package entities

import "gorm.io/gorm"

type RelaScheduleMessage struct {
	gorm.Model
	ScheduleId  int
	MessageId   int
	MessageType int
	AgentId     int
}
