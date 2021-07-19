package entities

import "gorm.io/gorm"

type Schedules struct {
	gorm.Model
	CreateBy       string
	AgentId        int
	ScheduleTitle  string
	SendDateTime   string
	ScheduleStatus int
}
