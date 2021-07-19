package entities

import "gorm.io/gorm"

type ConsumersSegments struct {
	gorm.Model
	CreateBy    string
	AgentId     int
	SegmentName string
	Description string
}
