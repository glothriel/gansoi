package mock

import "time"

type timestamp time.Time

type Metric struct {
	Timestamp time.Time
	Value     interface{}
}
type SingleHostMetrics map[string]Metric

type Collector struct {
	MetricsByHost map[string]*SingleHostMetrics
}

func (c *Collector) Add(host string, name string, timestamp time.Time, value interface{}) {
	_, ok := c.MetricsByHost[host]
	// Mutex here
	if !ok {
		m := make(SingleHostMetrics)
		c.MetricsByHost[host] = &m
	}
	(*c.MetricsByHost[host])[name] = Metric{Timestamp: timestamp, Value: value}
}

func NewCollector() *Collector {
	c := &Collector{}
	c.MetricsByHost = make(map[string]*SingleHostMetrics)
	return c
}
