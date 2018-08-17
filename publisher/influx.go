package main

import (
	"fmt"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
	"github.com/wattx/wwg-gobot/model"
)

const (
	db = "wemos"

	measurementReading = "reading"
)

// Influx described InfluxDB client
type Influx struct {
	client client.Client
}

// newInflux creates new InfluxDB client
func newInflux(url string) (*Influx, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: url,
	})

	if err != nil {
		return nil, err
	}
	defer c.Close()

	if err := c.Close(); err != nil {
		return nil, err
	}

	return &Influx{client: c}, nil
}

// WriteReading writes points to InfluxDB
func (ix *Influx) WriteReading(r model.Reading) error {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})
	if err != nil {
		return fmt.Errorf("unable to create batch points: %s", err)
	}

	// Create a point and add to batch
	tags := map[string]string{
		"type": r.Name(),
	}
	fields := map[string]interface{}{
		"value": r.Data(),
	}

	pt, err := client.NewPoint(measurementReading, tags, fields, time.Now())
	if err != nil {
		return fmt.Errorf("unable to create new point: %s", err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := ix.client.Write(bp); err != nil {
		return fmt.Errorf("unable to write to influx: %s", err)
	}

	return nil
}
