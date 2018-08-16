package main

import (
	"log"
	"strconv"

	"github.com/namsral/flag"
	"github.com/nats-io/nats"
)

const (
	topicHumidity = "humidity"
	topicMotion   = "motion"
)

var (
	influxURL = flag.String("influx_url", "", "influxdb url")
	natsURL   = flag.String("nats_url", "localhost:4222", "nats URL")
)

// byteToFloat32 converts slice of byte to float32
func byteToFloat32(b []byte) (float32, error) {
	value, err := strconv.ParseFloat(string(b), 32)
	if err != nil {
		return 0.0, err
	}

	h := float32(value)

	return h, nil
}

func byteToInt(b []byte) (int, error) {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func main() {
	flag.Parse()

	ix, err := newInflux(*influxURL)
	if err != nil {
		log.Fatal(err)
	}

	nc, err := newNATS(*natsURL)
	if err != nil {
		log.Fatal(err)
	}

	nc.Subscribe(topicHumidity, func(msg *nats.Msg) {
		h, err := byteToFloat32(msg.Data)
		if err != nil {
			log.Printf("unable to convert msg payload to float: %s", err)
		}

		log.Printf("received humidity: %.2f", h)

		err = ix.WriteHumidity(h)
		if err != nil {
			log.Printf("unable to publish humidity: %s", err)

			return
		}
	})

	nc.Subscribe(topicMotion, func(msg *nats.Msg) {
		i, err := byteToInt(msg.Data)
		if err != nil {
			log.Printf("unable to convert msg payload to float: %s", err)
		}

		log.Printf("received motion: %d", i)

		err = ix.WriteMotion(i)
		if err != nil {
			log.Printf("unable to publish motion: %s", err)

			return
		}
	})

	log.Println("Publisher is running.")

	<-make(chan int)
}
