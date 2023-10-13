package consumer

import (
	"log"

	"github.com/Shopify/sarama"
)

const (
	// KafkaConnectionString host:port
	KafkaConnectionString = "localhost:9092"

	// KafkaTopic topic name
	KafkaTopic = "test-topic"
)

func Consume() {
	// create consumer
	consumer, err := sarama.NewConsumer([]string{KafkaConnectionString}, nil)

	//check if connected
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to %s", KafkaConnectionString)

	// close connection
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// start consuming given topic
	partitionConsumer, err := consumer.ConsumePartition(KafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}

	// close connection
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// read messages sent to topic
	consumed := 0
	for {
		msg := <-partitionConsumer.Messages()
		// print message offset, (value, payload)
		log.Printf("Consumed message offset %d: %s:%s", msg.Offset, msg.Key, msg.Value)
		consumed++
	}

	// print number of consumed message - we will not reach this
	log.Printf("Consumed: %d", consumed)
	log.Print("Done")
}
