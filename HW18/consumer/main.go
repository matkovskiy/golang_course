package main

import (
	"encoding/json"
	"fmt"
	"log"
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

var small, medium, big int

func main() {
	go printMessage()
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	var forever chan struct{}
	res := Fruit{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			json.Unmarshal(d.Body, &res)

			switch size := res.Size; {
			case size < 40:
				log.Printf("It's small")
				small++
			case size > 60:
				log.Printf("It's big")
				big++
			default:
				log.Printf("It's medium")
				medium++
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func printMessage() {
	for {
		fmt.Println("Small: ", small)
		fmt.Println("Medium: ", medium)
		fmt.Println("Big: ", big)
		time.Sleep(1 * time.Minute)
	}
}
