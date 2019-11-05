package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	go server()
	go client()

	var a string
	fmt.Scanln(&a)
}

func server() {
	conn, ch, queue := getQueue()
	defer conn.Close()
	defer ch.Close()

	for count := 1; count <= 10; count++ {
		msg := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("Hello RabbitMQ %d", count)),
		}

		ch.Publish("", queue.Name, false, false, msg)
	}

}

func client() {
	conn, ch, queue := getQueue()
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	for msg := range msgs {
		log.Printf("Recieved message with message: %s", msg.Body)
	}
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	queue, err := ch.QueueDeclare("HelloChannel", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	return conn, ch, &queue
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
