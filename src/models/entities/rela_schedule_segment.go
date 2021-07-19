package entities

import "gorm.io/gorm"

type RelaScheduleSegment struct {
	gorm.Model
	ScheduleId int
	SegmentId  string
	AgentId    int
}
