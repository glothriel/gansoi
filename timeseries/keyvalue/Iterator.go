package keyvalue

import (
	"github.com/dgryski/go-tsz"
)

type Iterator interface {
	Err() error
	Next() bool
	Values() (uint32, float64)
}

type IteratorFactory interface {
	GetFor([]byte) (Iterator, error)
}

type GorillaIteratorFactory struct {
}

func (f *GorillaIteratorFactory) GetFor(b []byte) (Iterator, error) {
	return tsz.NewIterator(b)
}
