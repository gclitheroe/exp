// Use the Confluent Kafka platform https://docs.confluent.io/platform/current/platform-quickstart.html#quickstart
// wget https://raw.githubusercontent.com/confluentinc/cp-all-in-one/7.4.0-post/cp-all-in-one-kraft/docker-compose.yml
// docker-compose up -d
// visit http://localhost:9021/clusters
// Select the control center cluster
// Create a topic called quake using a protobuf schema with quake.proto schema from the protobuf dir in this repo.
// build and run this application
// go build
// ./quake-producer-kafka
package main

import (
	"flag"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/protobuf"
	"github.com/gclitheroe/exp/internal/quake"
	"log"
	"os"
)

type info struct {
	publicID string
	file     string
}

func main() {
	var bootstrap, topic, inDir, schemaRegistry string

	flag.StringVar(&bootstrap, "bootstrap", "localhost", "the Kafka bootstrap server")
	flag.StringVar(&schemaRegistry, "schema-registry", "http://localhost:8081", "url for the schema registry")
	flag.StringVar(&topic, "topic", "quake", "the topic to consume from")
	flag.StringVar(&inDir, "input-dir", "/work/quake", "directory with input quake protobuf files")

	flag.Parse()

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrap})
	if err != nil {
		log.Fatal(err)
	}

	defer p.Close()

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(schemaRegistry))
	if err != nil {
		log.Fatal(err)
	}

	serValue, err := protobuf.NewSerializer(client, serde.ValueSerde, protobuf.NewSerializerConfig())
	if err != nil {
		log.Fatal(err)
	}

	serKey, err := protobuf.NewSerializer(client, serde.KeySerde, protobuf.NewSerializerConfig())
	if err != nil {
		log.Fatal(err)
	}

	// Delivery report handler for produced messages.
	go func() {
		var i info
		var ok bool

		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v.", ev.TopicPartition)
				} else {
					i, ok = ev.Opaque.(info)
					if ok {
						log.Printf("Delivered message for %s", i.publicID)
						// the input file for the quake could be deleted here.
					}
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)

	files, err := os.ReadDir(inDir)
	if err != nil {
		log.Fatal(err)
	}

	var file string
	var k []byte
	var v []byte
	var q quake.Quake
	var key quake.Key

	for _, f := range files {
		file = inDir + string(os.PathSeparator) + f.Name()

		q, err = quake.Read(file)
		if err != nil {
			log.Println(err)
			// errored files could be moved or deleted here in a real application.
			continue
		}

		v, err = serValue.Serialize(topic, &q)
		if err != nil {
			log.Println(err)
			// errored files could be moved or deleted here in a real application.
			continue
		}

		key.QuakeID = q.PublicID

		k, err = serKey.Serialize(topic, &key)
		if err != nil {
			log.Println(err)
			// errored files could be moved or deleted here in a real application.
			continue
		}

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            k,
			Value:          v,
			Opaque:         info{file: file, publicID: q.PublicID},
		}, nil)
		if err != nil {
			log.Println(err)
		}
	}

	// Wait for message deliveries before shutting down
	for p.Flush(10000) > 0 {
		log.Println("waiting to flush outstanding messages.")
	}

}
