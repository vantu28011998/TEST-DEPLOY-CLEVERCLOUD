package database

import (
	"DialogOne/models/entities"
	"log"
	"gorm.io/gorm"
)
func CreateTable(db *gorm.DB, err error) {
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(
		&entities.SystemAgents{},
		&entities.SystemAccounts{},
		&entities.SystemBots{},
		&entities.ConsumersLineAccounts{},
		&entities.ConsumersSegments{},
		&entities.MessagesText{},
		&entities.Schedules{},
		&entities.RelaScheduleMessage{},
		&entities.RelaScheduleSegment{},
	)
}

