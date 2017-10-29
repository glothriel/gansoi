package keyvalue

import (
	"github.com/gansoi/gansoi/timeseries"
)

type Database struct {
	storage  KeyValueStorage
	factory  ItemCollectionFactory
	reporter timeseries.Reporter
}

func NewDatabase(storage KeyValueStorage, factory ItemCollectionFactory) *Database {
	return &Database{
		storage: storage,
		factory: factory,
		reporter: &Reporter{
			iteratorFactory: &GorillaIteratorFactory{},
		},
	}
}

func (d *Database) Store(host, name string, metric timeseries.Metric) {
	key := getKey(host, name)
	metrics, _ := d.storage.Get(key)
	if metrics == nil {
		metrics = d.factory.InitializeWith(metric)
	}
	metrics.Push(metric.Timestamp, metric.Value)
	d.storage.Set(key, metrics)
}

func (d *Database) Reporter() timeseries.Reporter {
	return nil
}

func getKey(host, metricName string) string {
	return host + "::" + metricName
}
