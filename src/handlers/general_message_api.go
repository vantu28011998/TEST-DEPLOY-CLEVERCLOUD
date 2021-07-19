package handlers

import (
	"DialogOne/models/notifications"
	MessageService "DialogOne/services/message"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetGeneralMessage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		query := r.URL.Query()
		page, present := query["page"]
		notification := notifications.Notification{
			Notification: "Error",
			Action:       "Get text messages which have been pagination",
			Status:       -1,
			Message:      "Params 'page' and 'limit' not null ",
		}
		if !present || len(page) == 0 {
			json.NewEncoder(w).Encode(notification)
			return
		}
		limit, present := query["limit"]
		if !present || len(limit) == 0 {
			json.NewEncoder(w).Encode(notification)
			return
		}
		intPage, err := strconv.Atoi(page[0])
		if err != nil {
			json.NewEncoder(w).Encode(notifications.Notification{
				Notification: "Error",
				Action:       "Get text messages which have been pagination",
				Status:       -1,
				Message:      "Params 'page' and 'limit' incorrect ",
			})
		}
		intLimit, err := strconv.Atoi(limit[0])
		if err != nil {
			json.NewEncoder(w).Encode(notifications.Notification{
				Notification: "Error",
				Action:       "Get text messages which have been pagination",
				Status:       -1,
				Message:      "Params 'page' and 'limit' incorrect ",
			})
		}
		messages, err := MessageService.GetGeneralMessages(intPage, intLimit)
		if err != nil {
			json.NewEncoder(w).Encode(notifications.Notification{
				Notification: "Error",
				Action:       "Get text messages which have been pagination",
				Status:       -1,
				Message:      "Error database connection",
			})
		}
		json.NewEncoder(w).Encode(messages)
	}
}
func SearchGeneralMessage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		query := r.URL.Query()
		page, present := query["page"]
		notification := notifications.Notification{
			Notification: "Error",
			Action:       "Get text messages which have been pagination",
			Status:       -1,
			Message:      "Params 'page' and 'limit' not null ",
		}
		if !present || len(page) == 0 {
			json.NewEncoder(w).Encode(notification)
			return
		}
		limit, present := query["limit"]
		if !present || len(limit) == 0 {
			json.NewEncoder(w).Encode(notification)
			return
		}
		intPage, err := strconv.Atoi(page[0])
		if err != nil {
			json.NewEncoder(w).Encode(notifications.Notification{
				Notification: "Error",
				Action:       "Get text messages which have been pagination",
				Status:       -1,
				Message:      "Params 'page' and 'limit' incorrect ",
			})
		}
		intLimit, err := strconv.Atoi(limit[0])
		if err != nil {
			json.NewEncoder(w).Encode(notifications.Notification{
				Notification: "Error",
				Action:       "Get text messages which have been pagination",
				Status:       -1,
				Message:      "Params 'page' and 'limit' incorrect ",
			})
		}
		types, present := query["type"]
		title, present := query["title"]
		MessageService.SearchGeneralMessages(intPage, intLimit, types, title[0])

	}
}
