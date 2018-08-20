package main

import (
	"github.com/namsral/flag"

	"gobot.io/x/gobot/platforms/firmata"
)

var (
	firmataURL = flag.String("firmata_url", "", "firmata TCP address")
)

func main() {
	flag.Parse()

	// new adaptors
	firmataAdaptor := firmata.NewTCPAdaptor(*firmataURL)

	// new wemos
	wemos := NewWemos(firmataAdaptor)
	wemos.Start()
}

