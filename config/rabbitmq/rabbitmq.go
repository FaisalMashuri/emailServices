package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/FaisalMashuri/emailServices/domain"
	"github.com/FaisalMashuri/emailServices/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQ(config models.AppConfig) *RabbitMQ {
	//fmt.Println(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.User, config.Password, config.Host, config.Port))
	conn, err := amqp.Dial(config.DSN_MQ)
	if err != nil {
		log.Fatal(err)

	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	queue, err := ch.QueueDeclare(
		"", // Update with your queue name
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}
	log.Println("RabbitMQ Connected")
	return &RabbitMQ{
		conn:    conn,
		channel: ch,
		queue:   queue,
	}
}

func (r *RabbitMQ) EmailConsumer(emailService domain.EmailDomain) {
	err := r.channel.ExchangeDeclare("user", "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = r.channel.QueueBind(
		r.queue.Name,    // queue name
		"user.register", // routing key
		"user",          // exchange
		false,
		nil)
	if err != nil {
		log.Println(fmt.Sprintf("Failed To register queue binding : %s", err.Error()))
	}
	msgs, err := r.channel.Consume(
		"",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println(fmt.Sprintf("Failed to register a consumer: %s", err.Error()))
	}

	go func() {
		for d := range msgs {
			// Process the received message, e.g., send an email
			fmt.Println("EXCHANGE : ", d.RoutingKey)
			switch d.RoutingKey {
			case "user.register":
				var eventEmitter models.EventEmitter
				var emailEmitter models.EmailEmitter
				fmt.Println("Body : ", d.Body)
				err := json.Unmarshal(d.Body, &eventEmitter)
				payload, _ := json.Marshal(&eventEmitter.Payload)
				err = json.Unmarshal(payload, &emailEmitter)
				log.Println("Succes Unmarshal Event: ", eventEmitter)
				log.Println("Succes Unmarshal Email: ", emailEmitter)

				err = emailService.SendEmail(emailEmitter.Email, emailEmitter.Otp)
				if err != nil {
					log.Println("Failed to send")
				}
			}
		}
	}()
}
func (r *RabbitMQ) Stop() {
	if err := r.channel.Close(); err != nil {
		log.Println("Failed to close channel:", err)
	}

	if err := r.conn.Close(); err != nil {
		log.Println("Failed to close connection:", err)
	}
}
