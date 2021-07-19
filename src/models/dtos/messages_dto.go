package dtos

type MessagesDto struct {
	Pagination Pagination        `json:"pagination"`
	Messages   []MessagesTextDto `json:"messages"`
}
