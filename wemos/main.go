package main

import (
	"github.com/namsral/flag"

	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot/platforms/nats"
)

var (
	firmataURL = flag.String("firmata_url", "", "firmata TCP address")
	natsURL    = flag.String("nats_url", "localhost:4222", "nats URL")
	natsTOPIC   = flag.String("nats_topic", "", "nats TOPIC")
)

func main() {
	flag.Parse()

	// new adaptors
	firmataAdaptor := firmata.NewTCPAdaptor(*firmataURL)
	natsAdaptor := nats.NewAdaptor(*natsURL, 23)

	// new wemos
	wemos := NewWemos(firmataAdaptor, natsAdaptor)
	wemos.Start()
}
