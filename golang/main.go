package main

import (
	"fmt"

	"github.com/dveselka/kafka-consumer/goalng/consumer"
)

func main() {
	// send message to Kafka
	message := consumer.Consume()
	fmt.Println(message)
}
