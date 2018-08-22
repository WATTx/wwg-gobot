package main

import (
	"fmt"
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot/platforms/nats"
)

const (
	topicHumidity = "humidity"
	topicMotion   = "motion"
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
	led := gpio.NewLedDriver(f, "14")
	bme := i2c.NewBME280Driver(f, i2c.WithBus(0), i2c.WithAddress(0x76))
	motion := gpio.NewPIRMotionDriver(f, "12")

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
		log.Printf("Humidity is %f", h)
		w.nats.Publish(topicHumidity, []byte(fmt.Sprintf("%.2f", h)))
	})

	w.motion.On(gpio.MotionDetected, func(data interface{}) {
		w.led.Off()
		log.Println("Motion detected")
		w.nats.Publish(topicMotion, []byte("1"))
	})

	w.motion.On(gpio.MotionStopped, func(data interface{}) {
		w.led.On()
		log.Printf("Motion stopped")
		w.nats.Publish(topicMotion, []byte("0"))
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
