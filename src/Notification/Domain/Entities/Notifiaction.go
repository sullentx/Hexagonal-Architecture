package entities

type Notification struct {
	ID       int    `json:"id"`
	ClientID int    `json:"client_id"`
	Content  string `json:"notification_content"`
}

func NewNotification(clientID int, content string) *Notification {
	return &Notification{ClientID: clientID, Content: content}
}
