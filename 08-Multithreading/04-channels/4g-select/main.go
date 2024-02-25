package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{id: i, Msg: "Hello, RabbitMQ!"}
			time.Sleep(1 * time.Second)
			c1 <- msg
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{id: i, Msg: "Hello, Kafka!"}
			time.Sleep(2 * time.Second)
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from RabbitMQ ID: %d - %s\n", msg.id, msg.Msg) //rabbitmq

		case msg := <-c2:
			fmt.Printf("Received from Kafka ID: %d - %s\n", msg.id, msg.Msg) //kafka

		case <-time.After(3 * time.Second):
			println("timeout")

			// default:
			//	println("default")
		}
	}
}
