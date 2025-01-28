package entities

type Notification struct {
	ID      int
	Content string
}

func NewNotification(Content string) *Notification {
	return &Notification{Content: Content}
}
