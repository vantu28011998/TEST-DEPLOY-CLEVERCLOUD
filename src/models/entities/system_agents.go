package entities

import "gorm.io/gorm"

type SystemAgents struct {
	gorm.Model
	BussinessName string
}
