package boltdb

import (
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

func (s *BoltDBStorage) Set(key string, collection ItemCollection) error {
	return s.bucket.Put([]byte(key), collection.Bytes())
}

func (s *BoltDBStorage) Get(key string) (ItemCollection, error) {
	collectionAsBytes := s.bucket.Get([]byte(key))
	return s.factory.FromBytes(collectionAsBytes)
}
