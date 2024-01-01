package main

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"strconv"
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
	for i := 0; i < 10; i++ {
		message := amqp091.Publishing{
			Headers: amqp091.Table{
				"sample": "header",
			},
			Body: []byte("Hello World " + strconv.Itoa(i)),
		}
		err = channel.PublishWithContext(ctx, "notification", "email", false, false, message)
		if err != nil {
			panic("error publishing a message to the queue:" + err.Error())
		}
	}
}
