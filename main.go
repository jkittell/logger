package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func main() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		log.Fatal(err, " ", amqpServerURL)
	}
	defer connectRabbitMQ.Close()

	ch, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"LoggerService", // queue name
		true,            // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	messages, err := ch.Consume(
		"LoggerService", // queue name
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no local
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	loggerService, err := newLogger()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	go func() {
		for message := range messages {
			var entry LogEntry
			err = json.Unmarshal(message.Body, &entry)
			if err != nil {
				log.Println(err)
			}
			loggerService.log(entry)
		}
	}()

	<-done
}
