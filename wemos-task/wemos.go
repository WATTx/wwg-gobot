package main

import (
	"log"
	_ "time"

	"gobot.io/x/gobot"
	_ "gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

// Board pin mapping to ESP8266 GPIO pin numbers
const  D0  = "16"
const  D1  = "5"
const  D2  = "4"
const  D3  = "0"
const  D4  = "2"
const  D5  = "14"
const  D6  = "12"
const  D7  = "13"
const  D8  = "15"
const  D9  = "3"
const  D10 = "1"
const LED_BUILTIN = "2"

// Wemos describes esp
type Wemos struct {
	firmata *firmata.TCPAdaptor
}

// NewWemos constructs new struct
func NewWemos(f *firmata.TCPAdaptor) *Wemos {
	return &Wemos{
		firmata: f,
	}
}

// Start runs robot
func (w *Wemos) Start() {
	robot := gobot.NewRobot(
		"bot",
		[]gobot.Connection{w.firmata},
		w.work,
	)

	robot.Start()
}

func (w *Wemos) work() {
	log.Println("Robot starts working...")
}
