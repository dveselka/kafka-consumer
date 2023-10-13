package main

import (
	"fmt"

	"consumer"
)

func main() {
	// send message to Kafka
	message := consumer.Produce()
	fmt.Println(message)
}
