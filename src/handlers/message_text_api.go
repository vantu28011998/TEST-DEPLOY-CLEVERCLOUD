package handlers

import (
	"DialogOne/models/entities"
	"DialogOne/models/notifications"
	"DialogOne/security"
	MessageService "DialogOne/services/message"
	"encoding/json"
	"log"
	"net/http"
)

func PostTxtMessages(w http.ResponseWriter, r *http.Request) {
	security.SetEnableCors(&w)
	switch r.Method {
	case "POST":
		var messageText entities.MessagesText
		var notification notifications.Notification
		err := json.NewDecoder(r.Body).Decode(&messageText)
		if err != nil {
			log.Fatal(err)
			notification = notifications.Notification{
				Notification: "Error",
				Action:       "Save a text message",
				Status:       -1,
				Message:      "Error parse json",
			}
			json.NewEncoder(w).Encode(notification)
			return
		}
		if messageText.Content == "" && messageText.Title == "" {
			notification = notifications.Notification{
				Notification: "Error",
				Action:       "Save a text message",
				Status:       -1,
				Message:      "Error json mapping || Title and content not null together",
			}
			json.NewEncoder(w).Encode(notification)
			return
		}
		notification = MessageService.CreateTextMessage(messageText)
		json.NewEncoder(w).Encode(notification)
		return
	}
}
