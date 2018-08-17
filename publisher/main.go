package main

import (
	"log"

	"github.com/namsral/flag"
	"github.com/nats-io/nats"
	"github.com/wattx/wwg-gobot/model"
)

var (
	influxURL = flag.String("influx_url", "", "influxdb url")
	natsURL   = flag.String("nats_url", "localhost:4222", "nats URL")
	natsTOPIC = flag.String("nats_topic", "", "nats TOPIC")
)

func newReadingProcessor(ix *Influx) chan<- []byte {
	ch := make(chan []byte)

	go func() {
		for payload := range ch {
			reading, err := model.UnpackEnvelope(payload)
			if err != nil {
				log.Printf("fail to unpack envelope `%s`, err: %s", payload, err)
				continue
			}
			err = ix.WriteReading(reading)

			if err != nil {
				log.Printf("unable to publish reading %s, err: %s", reading, err)
				break
			}
		}
		close(ch)
	}()

	return ch
}

func main() {
	flag.Parse()

	if *natsTOPIC == "" {
		log.Fatal("please specify NATS topic")
	}

	ix, err := newInflux(*influxURL)
	if err != nil {
		log.Fatal(err)
	}

	nc, err := newNATS(*natsURL)
	if err != nil {
		log.Fatal(err)
	}

	ch := newReadingProcessor(ix)
	nc.Subscribe(*natsTOPIC, func(msg *nats.Msg) {
		payload := msg.Data
		ch <- payload
		log.Printf("received payload: %s", payload)
	})

	log.Println("Publisher is running.")

	<-make(chan int)
}
