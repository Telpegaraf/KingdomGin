package consumer

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"kingdom/model"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"
)

type RMQConsumerDatabase interface {
	CreateUserCode(code *model.UserCode) error
}

type RMQConsumer struct {
	Queue      string
	Channel    *amqp091.Channel
	Connection *amqp091.Connection
	DB         RMQConsumerDatabase
}

func New(connectionString string) (rmg *RMQConsumer, err error) {
	log.Printf("Connecting to RabbitMQ at %s", connectionString)
	conn, err := amqp091.Dial(connectionString)
	if connectionString == "" {
		log.Fatal("RMQ_URL is not set")
	}

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

	rmq := &RMQConsumer{Queue: queue.Name, Channel: ch, Connection: conn}
	go func() { rmq.HandleMessage() }()

	return rmq, nil
}

func (r *RMQConsumer) HandleMessage() {
	msgs, err := r.Channel.Consume(
		r.Queue, "", true, false, false, false, nil)
	if err != nil {
		log.Println("Failed to register a consumer")
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {

			log.Printf("Received a message, email: %s", string(d.Body))
			if err != nil {
				log.Println("Failed to unmarshal the message", err)
				continue
			}
			r.SendEmail(string(d.Body))
		}
	}()
	log.Printf(" [*] Waiting for User messages.")
	<-forever
}

func (r *RMQConsumer) Publish(email string) {

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
		amqp091.Publishing{ContentType: "text/plain", Body: []byte(email)},
	)
	if err != nil {
		log.Println("Failed to publish a message")
		return
	}
	log.Println("Message published successfully")
}

func (r *RMQConsumer) SendEmail(email string) {
	code := GenerateCode()

	userCode := &model.UserCode{
		Code:  code,
		Email: email,
	}

	log.Printf("Sending email to %s with code %s", email, code)
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	message := "From: " + emailFrom + "\n" +
		"To: " + email + "\n" +
		"Subject: " + "Kingdom Register" + "\n" +
		"Your code for register is " + code + "\n"

	auth := smtp.PlainAuth("", emailFrom, emailPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, emailFrom, []string{email}, []byte(message))
	if err != nil {
		log.Println("Failed to send email:", err)
	} else {
		log.Println("Email sent successfully")
	}

	err = r.DB.CreateUserCode(userCode)
	if err != nil {
		log.Println("Failed to create a user code")
	}
}

func GenerateCode() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := ""
	for i := 0; i < 6; i++ {
		digit := rng.Intn(9) + 1
		code += fmt.Sprintf("%d", digit)
	}
	return code
}
