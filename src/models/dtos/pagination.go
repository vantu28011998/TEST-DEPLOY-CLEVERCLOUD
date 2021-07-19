package dtos

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	SumOfPage int `json:"sumOfPage"`
}
