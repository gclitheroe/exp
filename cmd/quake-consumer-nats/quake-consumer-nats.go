package main

import (
	"errors"
	"flag"
	"github.com/gclitheroe/exp/internal/quake"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var url, subject string

	flag.StringVar(&url, "url", nats.DefaultURL, "the NATS url")
	flag.StringVar(&subject, "subject", "quake", "the NATS subject to send messages to")

	flag.Parse()

	nc, err := nats.Connect(url)
	if err != nil {
		log.Panic(err)
	}

	defer nc.Close()

	sub, err := nc.SubscribeSync(subject)
	if err != nil {
		log.Panic(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	var q quake.Quake
	var m *nats.Msg

	for run {
		select {
		case _ = <-sigchan:
			log.Println("shutting down.")
			run = false
		default:
			m, err = sub.NextMsg(time.Second)
			if err != nil {
				if errors.Is(err, nats.ErrTimeout) {
					continue
				}
				log.Println(err)
				continue
			}

			err = proto.Unmarshal(m.Data, &q)
			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf("Received message from %s for %s\n", subject, q.PublicID)
		}
	}
}
