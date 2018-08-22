package main

import (
	"fmt"
	"log"
	"time"

	"github.com/0xAX/notificator"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/firmata"
)

// Board pin mapping to ESP8266 GPIO pin numbers
const D0 = "16"
const D1 = "5"
const D2 = "4"
const D3 = "0"
const D4 = "2"
const D5 = "14"
const D6 = "12"
const D7 = "13"
const D8 = "15"
const D9 = "3"
const D10 = "1"
const LED_BUILTIN = "2"

// Wemos describes esp
type Wemos struct {
	firmata *firmata.TCPAdaptor
	ledExt  *gpio.LedDriver
	motion  *gpio.PIRMotionDriver
	bme     *i2c.BME280Driver
}

var (
	notify *notificator.Notificator
)

// NewWemos constructs new struct
func NewWemos(f *firmata.TCPAdaptor) *Wemos {
	notify = notificator.New(notificator.Options{
		AppName: "Smart IoT",
	})

	ledExt := gpio.NewLedDriver(f, D5)
	motion := gpio.NewPIRMotionDriver(f, D6)
	bme := i2c.NewBME280Driver(f, i2c.WithBus(0), i2c.WithAddress(0x76))

	return &Wemos{
		firmata: f,
		ledExt:  ledExt,
		motion:  motion,
		bme:     bme,
	}
}

// Start runs robot
func (w *Wemos) Start() {
	robot := gobot.NewRobot(
		"bot",
		[]gobot.Connection{w.firmata},
		[]gobot.Device{w.ledExt, w.motion, w.bme},
		w.work,
	)

	robot.Start()
}

func (w *Wemos) work() {
	log.Println("Robot starts working...")

	gobot.Every(1*time.Second, func() {
		w.humidity()
	})

	w.motion.On(gpio.MotionDetected, func(s interface{}) {
		fmt.Println("motion detected")
		notify.Push("motion", "detected", "", notificator.UR_CRITICAL)
		w.ledExt.Off()
	})

	w.motion.On(gpio.MotionStopped, func(s interface{}) {
		fmt.Println("motion stopped")
		notify.Push("motion", "stopped", "", notificator.UR_NORMAL)
		w.ledExt.On()
	})
}

func (w *Wemos) toggleLed() {
	w.ledExt.Toggle()
}

// Humidity returns humidty value from bme280 sensor
func (w *Wemos) humidity() (float32, error) {
	h, err := w.bme.Humidity()
	log.Printf("%f", h)
	if err != nil {
		return 0.0, err
	}

	return h, nil
}
