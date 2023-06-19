// See the instructions in quake-producer-kafka for starting the broker and creating a topic
package main

import (
	"flag"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/protobuf"
	"github.com/gclitheroe/exp/internal/quake"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	var bootstrap, topic, group, schemaRegistry string

	flag.StringVar(&bootstrap, "bootstrap", "localhost", "the Kafka bootstrap server")
	flag.StringVar(&schemaRegistry, "schema-registry", "http://localhost:8081", "url for the schema registry")
	flag.StringVar(&topic, "topic", "quake", "the topic to consume from")
	flag.StringVar(&group, "group", "quakeConsumer", "the group")

	flag.Parse()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrap,
		"group.id":          group,
		// Start reading from the first message of each assigned
		// partition if there are no previously committed offsets
		// for this group.
		"auto.offset.reset": "earliest",
		// Do not automatically store offsets.
		// To enable at least once processing.
		"enable.auto.offset.store": false,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(schemaRegistry))
	if err != nil {
		log.Fatal(err)
	}

	ser, err := protobuf.NewDeserializer(client, serde.ValueSerde, protobuf.NewDeserializerConfig())
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{topic}, nil)

	run := true
	var q quake.Quake
	var msg *kafka.Message

	for run {
		select {
		case _ = <-sigchan:
			log.Println("shutting down.")
			run = false
		default:
			msg, err = c.ReadMessage(time.Second)
			if err != nil {
				// The client will automatically try to recover from all errors.
				// Timeout is not considered an error because it is raised by
				// ReadMessage in absence of messages.
				if !err.(kafka.Error).IsTimeout() {
					log.Printf("Consumer error: %v (%v)\n", err, msg)
				}
				continue
			}

			err = ser.DeserializeInto(topic, msg.Value, &q)
			if err != nil {
				log.Println(err)
				continue
			}

			// Any processing happens here e.g., store in a DB
			log.Printf("Received message for %s", q.PublicID)

			// Once processing is complete store the offsets.
			// This ensures at least once processing.
			_, err = c.StoreMessage(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
