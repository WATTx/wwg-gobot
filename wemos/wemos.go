package main

import (
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
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

	led    *gpio.LedDriver
}

// NewWemos constructs new struct
func NewWemos(f *firmata.TCPAdaptor) *Wemos {
	// new drivers
	led := gpio.NewLedDriver(f, D4)
	return &Wemos{
		firmata: f,
		led:     led,
	}
}

// Start runs robot
func (w *Wemos) Start() {
	robot := gobot.NewRobot(
		"bot",
		[]gobot.Connection{w.firmata},
		[]gobot.Device{w.led},
		w.work,
	)

	robot.Start()
}

func (w *Wemos) work() {
	log.Println("Robot starts working...")

	gobot.Every(1*time.Second, func() {
		w.led.Toggle()
	})
}

