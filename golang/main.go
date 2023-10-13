package main

import (
	"fmt"
)

func main() {
	// send message to Kafka
	message := consumer.Consume()
	fmt.Println(message)
}
