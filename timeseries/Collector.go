package timeseries

import "time"

type Collector interface {
	Add(host string, name string, timestamp time.Time, value interface{})
}
