package keyvalue

import "github.com/gansoi/gansoi/timeseries"

type Reporter struct {
	storage         KeyValueStorage
	iteratorFactory IteratorFactory
}

func (r *Reporter) PastMetrics(q timeseries.MetricsQuery) []timeseries.Metric {
	key := getKey(q.Host, q.Name)
	itemCollection, getError := r.storage.Get(key)
	if getError != nil {
		panic("asd")
	}
	iterator, factoryError := r.iteratorFactory.GetFor(itemCollection.Bytes())
	if factoryError != nil {
		panic("asd")
	}
	metrics := make([]timeseries.Metric, 0)
	for iterator.Next() {
		timestamp, value := iterator.Values()
		metrics = append(metrics, timeseries.Metric{
			Timestamp: timestamp,
			Value:     value,
		})
	}
	return metrics
}

func (r *Reporter) LiveMetrics(q timeseries.MetricsQuery) timeseries.LiveMetrics {
	return nil
}
