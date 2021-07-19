package entities

import "gorm.io/gorm"

type SystemAccounts struct {
	gorm.Model
	AgentId  int
	Username string
	Email    string
	Password string
	
}
