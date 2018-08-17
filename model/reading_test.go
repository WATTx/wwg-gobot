package model

import (
	"testing"
)

func TestEncodeEnvelope(t *testing.T) {
	en := NewEnvelope(Humidity, &HumidityReading{51.2})
	excepted := `{"Type":0,"Data":{"Value":51.2}}`
	if string(en.Encode()) != excepted {
		t.Errorf("not equal")
	}
}
func TestDecodeEnvelope(t *testing.T) {
	payload := []byte(`{"Type":0,"Data":{"Value":51.2}}`)
	reading, err := UnpackEnvelope(payload)
	if err != nil {
		t.Errorf("fail to unpack err: %s", err)
	}
	excepted := HumidityReading{51.2}
	if reading.Name() != excepted.Name() {
		t.Errorf("not equal, reading: %s, excepted: %s", reading.Name(), excepted.Name())
	} else if reading.Data() != excepted.Data() {
		t.Errorf("not equal, reading: %s, excepted: %s", reading.Data(), excepted.Data())
	}
}
