package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	var qName = "DemoRBMQ"
	fmt.Println("Consumer application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	defer ch.Close()

	messages, err := ch.Consume(
		qName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range messages {
			fmt.Printf("Received a message: %s\n", d.Body)
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
