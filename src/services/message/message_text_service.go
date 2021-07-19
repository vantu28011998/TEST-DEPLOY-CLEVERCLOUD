package message

import (
	"DialogOne/clients/postgresql"
	"DialogOne/models/entities"
	"DialogOne/models/notifications"
	"log"
)

func CreateTextMessage(messagesText entities.MessagesText) notifications.Notification {
	db, err := postgresql.GetConnection()
	if err != nil {
		log.Fatal(err)
		return notifications.Notification{
			Notification: "Error",
			Action:       "Save a text message",
			Status:       -1,
			Message:      "Error database connection",
		}
	}
	db.Create(&messagesText)
	return notifications.Notification{
		Notification: "Success",
		Action:       "Save a text message",
		Status:       1,
		Message:      "Save a text message successfully",
	}
}


