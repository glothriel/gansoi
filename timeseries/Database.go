package timeseries

// Database serves as an quick write-through cache for the check data
type Database interface {
	Collector() Collector
	Reporter() Reporter
}
