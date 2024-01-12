package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
)

func main() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		log.Fatal(err, " ", amqpServerURL)
	}
	defer connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	ch, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// With the instance and declare Queues that we can
	// publish and subscribe to.
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

	// Subscribing to LogService for getting messages.
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
		log.Println(err)
	}

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	logger, err := NewLogger()
	if err != nil {
		log.Fatal(err)
	}
	// Make a channel to receive messages into infinite loop.
	done := make(chan bool)

	go func() {
		for message := range messages {
			logger.WriteLog(message.Body)
		}
	}()

	<-done
}
