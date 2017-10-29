package keyvalue

import (
	"errors"
	"fmt"

	"github.com/boltdb/bolt"
)

type KeyValueStorage interface {
	Get(string) (ItemCollection, error)
	Set(string, ItemCollection) error
}

type BoltDBStorage struct {
	bucket  bolt.Bucket
	factory ItemCollectionFactory
}

type BoldDbSavedMetric struct {
	ID     string
	Metric []byte
}

func (o *BoldDbSavedMetric) SetID() {

}

func (o *BoldDbSavedMetric) GetID() string {
	return o.ID
}

func NewBoltDBStorage(bucket bolt.Bucket, factory ItemCollectionFactory) *BoltDBStorage {
	s := &BoltDBStorage{}
	s.bucket = bucket
	s.factory = factory
	return s
}

func (s *BoltDBStorage) Set(key string, collection ItemCollection) error {
	fmt.Println("Saving something to BoldDB")
	return s.bucket.Put([]byte(key), collection.Bytes())
}

func (s *BoltDBStorage) Get(key string) (ItemCollection, error) {
	collectionAsBytes := s.bucket.Get([]byte(key))
	return s.factory.FromBytes(collectionAsBytes)
}

type MapInMemStorage struct {
	data    map[string][]byte
	factory ItemCollectionFactory
}

func NewInMemStorage(factory ItemCollectionFactory) *MapInMemStorage {
	s := &MapInMemStorage{}
	s.data = make(map[string][]byte)
	s.factory = factory
	return s
}

func (s *MapInMemStorage) Set(key string, collection ItemCollection) error {
	fmt.Println("Saving " + key)
	s.data[key] = collection.Bytes()
	return nil
}

func (s *MapInMemStorage) Get(key string) (ItemCollection, error) {
	theItem, ok := s.data[key]
	if !ok {
		return nil, errors.New("Could not find " + key)
	}
	return s.factory.FromBytes(theItem)
}
