package memcacheex

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	_ Client = (*memcache.Client)(nil)
	_ Client = (*ClientWrapper)(nil)
)

type Client interface {
	FlushAll() error
	Get(key string) (item *memcache.Item, err error)
	Touch(key string, seconds int32) (err error)
	GetMulti(keys []string) (map[string]*memcache.Item, error)
	Set(item *memcache.Item) error
	Add(item *memcache.Item) error
	Replace(item *memcache.Item) error
	CompareAndSwap(item *memcache.Item) error
	Delete(key string) error
	DeleteAll() error
	Ping() error
	Increment(key string, delta uint64) (newValue uint64, err error)
	Decrement(key string, delta uint64) (newValue uint64, err error)
}
