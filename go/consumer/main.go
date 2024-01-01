package main

import (
	"context"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	connection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	ctx := context.Background()
	emailConsumer, err := channel.ConsumeWithContext(ctx, "email", "consumer", true, false, false, false, nil)
	if err != nil {
		panic("could not register consumer:" + err.Error())
	}

	for message := range emailConsumer {
		fmt.Println("routing key: ", message.RoutingKey)
		fmt.Println("message body: ", string(message.Body))
	}
}
