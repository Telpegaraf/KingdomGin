package consumer

import (
	"github.com/streadway/amqp"
	"log"
)

func Handler(queue string, msg amqp.Delivery, err error) {
	if err != nil {
		log.Printf("Error occurred in RMQ consumer")
	}
	log.Printf("Message received on '%s' queue: %s", queue, string(msg.Body))
}
