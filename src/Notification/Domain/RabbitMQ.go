package domain

type IMessagePublisher interface {
	PublishMessage(message string, id int, name string) error
}
