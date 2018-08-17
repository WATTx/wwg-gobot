package model

import (
	"encoding/json"
	"fmt"
)

// ReadingType reading type
type ReadingType int

const (
	// Humidity reading type
	Humidity ReadingType = iota
	// Motion reading type
	Motion
)

// Reading common reading functions
type Reading interface {
	Name() string
	Data() interface{}
	Decode(json.RawMessage) error
}

// HumidityReading humdity reading
type HumidityReading struct {
	Value float32
}

// MotionReading motion reading
type MotionReading struct {
	Value int
}

// Envelope encapsule reading with its type
type Envelope struct {
	Type ReadingType
	Data interface{}
}

// NewEnvelope create new envelope
func NewEnvelope(readingType ReadingType, reading Reading) *Envelope {
	return &Envelope{
		Type: readingType,
		Data: reading,
	}
}

// UnpackEnvelope unpack envelope to its original reading
func UnpackEnvelope(payload []byte) (Reading, error) {
	var data json.RawMessage
	env := Envelope{
		Data: &data,
	}

	err := json.Unmarshal(payload, &env)
	if err != nil {
		return nil, err
	}

	var reading Reading
	switch env.Type {
	case Humidity:
		reading = &HumidityReading{}
	case Motion:
		reading = &MotionReading{}
	}
	if err := reading.Decode(data); err != nil {
		return nil, fmt.Errorf("unable to unpack reading, err: %s", err)
	}
	return reading, nil
}

// Encode encode envelope to []byte
func (e Envelope) Encode() []byte {
	buf, _ := json.Marshal(e)
	return buf
}

// Name name of the humidity reading
func (hr HumidityReading) Name() string {
	return "humidity"
}

// Data data of reading
func (hr HumidityReading) Data() interface{} {
	return hr.Value
}

// Decode decode raw message to humdidity reading
func (hr *HumidityReading) Decode(raw json.RawMessage) error {
	msg := new(HumidityReading)
	err := json.Unmarshal(raw, &msg)
	if err != nil {
		return err
	}
	hr.Value = msg.Value
	return nil
}

func (hr HumidityReading) String() string {
	return fmt.Sprintf("<HumidityReading HumidityReading=%.2f>", hr.Value)
}

// Name name of the motion reading
func (mr MotionReading) Name() string {
	return "motion"
}

// Data of reading
func (mr MotionReading) Data() interface{} {
	return mr.Value
}

// Decode decode raw message to motion reading
func (mr *MotionReading) Decode(raw json.RawMessage) error {
	msg := new(MotionReading)
	err := json.Unmarshal(raw, &msg)
	if err != nil {
		return err
	}
	mr.Value = msg.Value
	return nil
}

func (mr MotionReading) String() string {
	return fmt.Sprintf("<MotionReading Value=%d>", mr.Value)
}
