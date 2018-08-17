package main

import (
	"log"
	"time"

	"github.com/wattx/wwg-gobot/model"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot/platforms/nats"
)

// Wemos describes esp
type Wemos struct {
	firmata *firmata.TCPAdaptor
	nats    *nats.Adaptor

	led    *gpio.LedDriver
	bme    *i2c.BME280Driver
	motion *gpio.PIRMotionDriver
}

// NewWemos constructs new struct
func NewWemos(f *firmata.TCPAdaptor, n *nats.Adaptor) *Wemos {
	// new drivers
	led := gpio.NewLedDriver(f, "2")
	bme := i2c.NewBME280Driver(f, i2c.WithBus(0), i2c.WithAddress(0x76))
	motion := gpio.NewPIRMotionDriver(f, "14")

	return &Wemos{
		firmata: f,
		nats:    n,
		led:     led,
		bme:     bme,
		motion:  motion,
	}
}

// Start runs robot
func (w *Wemos) Start() {
	robot := gobot.NewRobot(
		"bot",
		[]gobot.Connection{w.firmata, w.nats},
		[]gobot.Device{w.led, w.bme, w.motion},
		w.work,
	)

	robot.Start()
}

func (w *Wemos) work() {
	log.Println("Robot starts working...")

	gobot.Every(1*time.Second, func() {
		h, err := w.humidity()
		if err != nil {
			log.Printf("unable to get humidty: %s", err)
		}

		msg := model.NewEnvelope(model.Humidity, &model.HumidityReading{h})
		w.nats.Publish(*natsTOPIC, msg.Encode())
	})

	w.motion.On(gpio.MotionDetected, func(data interface{}) {
		w.led.Off()

		msg := model.NewEnvelope(model.Motion, &model.MotionReading{1})
		w.nats.Publish(*natsTOPIC, msg.Encode())
	})

	w.motion.On(gpio.MotionStopped, func(data interface{}) {
		w.led.On()

		msg := model.NewEnvelope(model.Motion, &model.MotionReading{0})
		w.nats.Publish(*natsTOPIC, msg.Encode())
	})
}

// Humidity returns humidty value from bme280 sensor
func (w *Wemos) humidity() (float32, error) {
	h, err := w.bme.Humidity()

	if err != nil {
		return 0.0, err
	}

	return h, nil
}

func (w *Wemos) toggleLED() error {
	return w.led.Toggle()
}
