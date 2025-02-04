package entities

type Notification struct {
	ID       int
	ClientID int
	Content  string
}

func NewNotification(clientID int, content string) *Notification {
	return &Notification{ClientID: clientID, Content: content}
}
