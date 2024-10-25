package consumer

import (
	"github.com/streadway/amqp"
	"log"
)

type RConsumer struct {
	Queue      string
	Channel    *amqp.Channel
	Connection *amqp.Connection
}

//type RMQConsumer struct {
//	Queue            string
//	ConnectionString string
//	MsgHandler       func(queue string, msg amqp.Delivery, err error)
//}
//
//func (r *RMQConsumer) OnError(err error, msg string) {
//	if err != nil {
//		r.MsgHandler(r.Queue, amqp.Delivery{}, err)
//	}
//}
//
//func (r *RMQConsumer) Consume() {
//	conn, err := amqp.Dial(r.ConnectionString)
//	r.OnError(err, "Failed to connect to RabbitMQ")
//	defer conn.Close()
//
//	ch, err := conn.Channel()
//	r.OnError(err, "Failed to open a channel")
//	defer ch.Close()
//
//	q, err := ch.QueueDeclare(
//		r.Queue,
//		false,
//		false,
//		false,
//		false,
//		nil)
//	r.OnError(err, "Failed to declare a queue")
//
//	msgs, err := ch.Consume(
//		q.Name, "", false, false, false, false, nil)
//	r.OnError(err, "Failed to register a consumer")
//
//	forever := make(chan bool)
//	go func() {
//		for d := range msgs {
//			r.MsgHandler(r.Queue, d, nil)
//		}
//	}()
//	log.Println("Started listening for messages on '%s' queue", q)
//	<-forever
//}

func New(connectionString string) (rmg *RConsumer, err error) {
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		log.Println("Failed to connect to RabbitMQ")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel")
		panic(err)
	}

	queue, err := ch.QueueDeclare(
		"UserQueue",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Println("Failed to declare a queue")
		ch.Close()
		conn.Close()
		panic(err)
	}

	//msgs, err := ch.Consume(
	//	queue.Name, "", false, false, false, false, nil)
	//if err != nil {
	//	log.Println("Failed to register a consumer")
	//	panic(err)
	//}
	//forever := make(chan bool)
	//go func() {
	//	for d := range msgs {
	//		log.Printf("Received a message: %s", d.Body)
	//	}
	//}()
	//log.Println("Started listening for messages on '%s' queue", queue)
	//go func() {
	//	<-forever
	//}()

	return &RConsumer{Queue: queue.Name, Channel: ch, Connection: conn}, nil
}

func (r *RConsumer) Publish(username string, email string) {
	message := "" + username + " " + email

	if r.Channel == nil {
		log.Println("Failed to connect to RabbitMQ")
		return
	}
	if r.Connection == nil {
		log.Println("Failed to connect to RabbitMQ - Connection is nil")
		return
	}

	err := r.Channel.Publish(
		"",
		r.Queue,
		false,
		false,
		amqp.Publishing{ContentType: "text/plain", Body: []byte(message)},
	)
	if err != nil {
		log.Println("Failed to publish a message")
		return
	}
	log.Println("Message published successfully")
}
