package message

import (
	"DialogOne/clients/postgresql"
	"DialogOne/models/dtos"
	"DialogOne/models/entities"
)

//GET TEXT MESSAGE
func GetGeneralMessages(page int, limit int) (dtos.MessagesDto, error) {
	db, err := postgresql.GetConnection()
	if err != nil {
		return dtos.MessagesDto{}, err
	}
	offset := (page - 1) * limit
	var messages []entities.MessagesText
	var messageDtos []dtos.MessagesTextDto
	db.Limit(limit).Offset(offset).Find(&messages)
	var count int64
	db.Table("messages_texts").Count(&count)
	sumOfPage := count/int64(limit) + 1
	for _, mapping := range messages {
		messageDto := dtos.MessagesTextDto{
			Id:       mapping.ID,
			Title:    mapping.Title,
			Content:  mapping.Content,
			Status:   mapping.Status,
			CreateAt: mapping.CreatedAt,
		}
		messageDtos = append(messageDtos, messageDto)
	}
	return dtos.MessagesDto{
		Pagination: dtos.Pagination{
			Page:      page,
			Limit:     limit,
			SumOfPage: int(sumOfPage),
		},
		Messages: messageDtos,
	}, nil
}
func SearchGeneralMessages(page int, limit int, types []string, title string) {
	if len(types) == 1 {
		if types[0] == "0" {
			//chua duoc tra ve de loc tiep
		}
		if types[0] == "1" {
 
		}
	}
	if len(types) == 2 {

	}
	if len(title) > 0 && len(types) == 1 {
		if types[0] == "0" {

		}
		if types[0] == "1" {

		}
	}
	if len(title) > 0 && len(types) == 2 {

	}
}
