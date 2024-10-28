package consumer

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func Handler(queue string, msg amqp091.Delivery, err error) {
	if err != nil {
		log.Printf("Error occurred in RMQ consumer")
	}
	log.Printf("Message received on '%s' queue: %s", queue, string(msg.Body))
}
