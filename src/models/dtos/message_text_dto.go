package dtos

import "time"

type MessagesTextDto struct {
	Id       uint     `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Status   int       `json:"status"`
	CreateAt time.Time `json:"createAt"`
}
