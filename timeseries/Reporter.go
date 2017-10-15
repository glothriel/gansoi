package timeseries

// LiveMetrics is a channel that streams live metrics as soon as they appear
type LiveMetrics chan *Metric

// Reporter returns all metrics that meet certain conditions
type Reporter interface {
	PastMetrics(MetricsQuery) []Metric
	LiveMetrics(MetricsQuery) LiveMetrics
}

// MetricsQuery gathers popular query attributes
type MetricsQuery struct {
	Host string
	// Name of the metric to search for
	Name string
	// Zero means that the query is not bound by start time
	TimeStart uint32
	// Zero means that the query is not bound by end time
	TimeEnd uint32
	// If empty or nil just dump all the metrics without granularity
	GranularityLevels []Granularity
}

// Granularity specifies how big amount of metrics in certain period of time we should report
type Granularity struct {
	OlderThanSeconds     uint16
	GranularityInSeconds uint16
}
