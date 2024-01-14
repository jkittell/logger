package main

import (
	"fmt"
	"github.com/wagslane/go-rabbitmq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	loggerService, err := newLogger()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := rabbitmq.NewConn(
		os.Getenv("RABBITMQ_SERVER_URL"),
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	consumer, err := rabbitmq.NewConsumer(
		conn,
		func(d rabbitmq.Delivery) rabbitmq.Action {
			loggerService.log(string(d.Body))
			return rabbitmq.Ack
		},
		"q.logs",
		rabbitmq.WithConsumerOptionsRoutingKey("logs"),
		rabbitmq.WithConsumerOptionsExchangeName("events"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	<-done
}
