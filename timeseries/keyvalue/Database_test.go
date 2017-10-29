package keyvalue

import (
	"testing"

	"github.com/gansoi/gansoi/timeseries"
)

type MockItemCollection struct {
}

func (c *MockItemCollection) Push(timestamp uint32, value float64) {

}

type MockItemCollectionFactory struct {
}

func (c *MockItemCollectionFactory) FromBytes([]byte) (ItemCollection, error) {
	return &MockItemCollection{}, nil
}

func (c *MockItemCollectionFactory) InitializeWith(timeseries.Metric) ItemCollection {
	return &MockItemCollection{}
}

func (c *MockItemCollection) Bytes() []byte {
	return []byte{}
}

type MockStorage struct {
}

func (s *MockStorage) Get(key string) (ItemCollection, error) {
	return &MockItemCollection{}, nil
}

func (s *MockStorage) Set(key string, value ItemCollection) error {
	return nil
}

func TestStore(t *testing.T) {
	d := &Database{
		storage: &MockStorage{},
		factory: &MockItemCollectionFactory{},
	}
	d.Store(
		"kitty.com",
		"food.KilosPerDay",
		timeseries.Metric{
			Timestamp: 1337,
			Value:     133.7,
		},
	)
}
