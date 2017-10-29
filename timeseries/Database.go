package timeseries

// Metric stores single gathered metric for single host
type Metric struct {
	Timestamp uint32
	Value     float64
}

// Database serves as an quick write-through cache for the check data
type Database interface {
	Store(host, name string, metric Metric)
	Reporter() Reporter
}
