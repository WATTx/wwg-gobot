package main

import (
	"fmt"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	db = "wemos"

	measurementHumidity = "humidity"
	measurementMotion   = "motion"
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

// WriteHumidity writes points to InfluxDB
func (ix *Influx) WriteHumidity(h float32) error {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})
	if err != nil {
		return fmt.Errorf("unable to create batch points: %s", err)
	}

	// Create a point and add to batch
	tags := map[string]string{}
	fields := map[string]interface{}{
		"value": h,
	}

	pt, err := client.NewPoint(measurementHumidity, tags, fields, time.Now())
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

// WriteMotion writes points to InfluxDB
func (ix *Influx) WriteMotion(i int) error {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})
	if err != nil {
		return fmt.Errorf("unable to create batch points: %s", err)
	}

	// Create a point and add to batch
	tags := map[string]string{}
	fields := map[string]interface{}{
		"value": i,
	}

	pt, err := client.NewPoint(measurementMotion, tags, fields, time.Now())
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
