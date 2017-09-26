package mock

type MockDatabase struct {
	collector *Collector
	reporter  *Reporter
}

func (d *MockDatabase) Collector() *Collector {
	return d.collector
}

func (d *MockDatabase) Reporter() *Reporter {
	return d.reporter
}
