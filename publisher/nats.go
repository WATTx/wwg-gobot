package main

import (
	"github.com/nats-io/nats"
)

// NATS describes NATS connection
type NATS struct {
	conn *nats.Conn
}

// newNATS creates new NATS connection
func newNATS(url string) (*NATS, error) {
	c, err := nats.Connect("nats://docker.test:4222")
	if err != nil {
		return nil, err
	}

	return &NATS{conn: c}, nil
}

// Subscribe exposes nats.Subscribe method
func (n *NATS) Subscribe(topic string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return n.conn.Subscribe(topic, cb)
}
