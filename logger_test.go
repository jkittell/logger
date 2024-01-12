package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
	"testing"
)

func TestLogger_WriteLog(t *testing.T) {
	// Create a new RabbitMQ connection.
	conn, err := amqp.Dial(os.Getenv("AMQP_SERVER_URL"))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	ch, err := conn.Channel()
	if err != nil {
		t.Fatal(err)
	}
	defer ch.Close()

	// With the instance and declare Queues that we can
	// publish and subscribe to.

	for i := 0; i < 10; i++ {
		body := fmt.Sprintf("hello %d", i)
		// Create a message to publish.
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		}

		if err = ch.Publish(
			"",              // exchange
			"LoggerService", // queue name
			false,           // mandatory
			false,           // immediate
			message,         // message to publish
		); err != nil {
			t.Fatal(err)
		}
	}
}
