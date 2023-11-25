package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type Fruit struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

func main() {
	conn, err := amqp.Dial("amqp://rmuser:rmpassword@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newApple := &Fruit{
		Name: "Apple",
		Size: rand.Intn(100)}
	newAppleBody, _ := json.Marshal(newApple)
	fmt.Println(string(newAppleBody))

	// body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(newAppleBody),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", newAppleBody)
}
