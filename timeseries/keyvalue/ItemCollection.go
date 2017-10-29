package keyvalue

import "github.com/gansoi/gansoi/timeseries"
import (
	"github.com/dgryski/go-tsz"
)

type ItemCollection interface {
	Push(timestamp uint32, value float64)
	Bytes() []byte
}

type ItemCollectionFactory interface {
	FromBytes([]byte) (ItemCollection, error)
	InitializeWith(timeseries.Metric) ItemCollection
}

type GorillaItemCollection tsz.Series

type GorillaItemCollectionFactory struct {
}

func (f *GorillaItemCollectionFactory) InitializeWith(m timeseries.Metric) ItemCollection {
	return tsz.New(m.Timestamp)
}

func (f *GorillaItemCollectionFactory) FromBytes(data []byte) (ItemCollection, error) {
	collection := &tsz.Series{}
	return collection, collection.UnmarshalBinary(data)
}
