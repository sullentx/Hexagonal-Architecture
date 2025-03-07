package adapters

import (
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn *amqp091.Connection
	ch   *amqp091.Channel
}

type NotificationMessage struct {
	ClientID   int    `json:"client_id"`
	ClientName string `json:"client_name"`
	Content    string `json:"notification_content"`
}

func InitRabbitMQ() *RabbitMQAdapter {
	conn, err := amqp091.Dial("amqp://uriel:eduardo117@3.228.81.226:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	_, err = ch.QueueDeclare(
		"notifications", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	return &RabbitMQAdapter{conn: conn, ch: ch}
}

func (r *RabbitMQAdapter) PublishMessage(message string, id int, name string) error {
	// Crear objeto de mensaje
	notificationMsg := NotificationMessage{
		ClientID:   id,
		ClientName: name,
		Content:    message,
	}

	// Convertir a JSON
	jsonData, err := json.Marshal(notificationMsg)
	if err != nil {
		log.Printf("Failed to marshal message to JSON: %v", err)
		return err
	}

	// Publicar el mensaje JSON
	err = r.ch.Publish(
		"",              // exchange
		"notifications", // routing key
		false,           // mandatory
		false,           // immediate
		amqp091.Publishing{
			ContentType: "application/json", // Cambiar a application/json
			Body:        jsonData,
		})
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
		return err
	}

	log.Printf(" [x] Sent JSON notification: %s", string(jsonData))
	return nil
}

func (r *RabbitMQAdapter) Close() {
	r.ch.Close()
	r.conn.Close()
}
