package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ with golang demo")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Successfully connected to local rabbitmq instance")

	var input string

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"DemoRBMQ", false, false, false, false, nil,
	)

	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	for {
		fmt.Scanf("%s", &input)
		if input == "exit" {
			break
		}
		fmt.Println("Sending message:", input)

		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(input),
			},
		)
		if err != nil {
			fmt.Println("Error:", err)
			panic(err)
		}
		fmt.Println("Message published: ", input)
	}

	fmt.Println("Exiting...")
}
