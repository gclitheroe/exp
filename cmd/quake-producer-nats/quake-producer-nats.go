package main

import (
	"flag"
	"github.com/gclitheroe/exp/internal/quake"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var url, subject, inDir string

	flag.StringVar(&url, "url", nats.DefaultURL, "the NATS url")
	flag.StringVar(&subject, "subject", "quake", "the NATS subject to send messages to")
	flag.StringVar(&inDir, "input-dir", "/work/quake", "directory with input quake protobuf files")

	flag.Parse()

	nc, err := nats.Connect(url)
	if err != nil {
		log.Panic(err)
	}

	defer nc.Drain()

	files, err := os.ReadDir(inDir)
	if err != nil {
		log.Fatal(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	var file string
	var q quake.Quake
	var b []byte

send:
	for _, f := range files {
		select {
		case _ = <-sigchan:
			log.Println("shutting down.")
			break send
		default:
			file = inDir + string(os.PathSeparator) + f.Name()

			q, err = quake.Read(file)
			if err != nil {
				log.Println(err)
				// errored files could be moved or deleted here in a real application.
				continue
			}

			b, err = proto.Marshal(&q)
			if err != nil {
				log.Println(err)
				// errored files could be moved or deleted here in a real application.
				continue
			}

			// With NATS delivery is at most once.
			err = nc.Publish(subject, b)
			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf("Published message to %s for %s\n", subject, q.PublicID)
			// a real application could delete the input file here.
		}
	}
}
