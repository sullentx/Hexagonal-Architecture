package adapters

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn *amqp091.Connection
	ch   *amqp091.Channel
}

func InitRabbitMQ() *RabbitMQAdapter {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
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

func (r *RabbitMQAdapter) PublishMessage(message string) error {
	err := r.ch.Publish(
		"",              // exchange
		"notifications", // routing key
		false,           // mandatory
		false,           // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
		return err
	}

	log.Printf(" [x] Sent %s", message)
	return nil
}

func (r *RabbitMQAdapter) Close() {
	r.ch.Close()
	r.conn.Close()
}
