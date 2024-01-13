package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"testing"
)

func TestLogger_WriteLog(t *testing.T) {
	conn, err := amqp.Dial(os.Getenv("AMQP_SERVER_URL"))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Fatal(err)
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

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("hello %d", i)
		entry := LogEntry{
			Name:    "testing",
			Message: msg,
		}

		data, _ := json.Marshal(entry)
		message := amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
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
