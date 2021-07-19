package notifications

type Notification struct {
	Notification string `json:"notification"`
	Action       string `json:"action"`
	Status       int    `json:"status"`
	Message      string `json:"message"`
}
