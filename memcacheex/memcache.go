package memcacheex

import (
	"github.com/bradfitz/gomemcache/memcache"
)

type ClientWrapper struct {
	client   Client
	registry *callbackRegistry
}

func NewClientWrapper(client Client) *ClientWrapper {
	return &ClientWrapper{
		client: client,
		registry: &callbackRegistry{
			flushAll:       &callbacks{},
			get:            &callbacks{},
			touch:          &callbacks{},
			getMulti:       &callbacks{},
			set:            &callbacks{},
			add:            &callbacks{},
			replace:        &callbacks{},
			compareAndSwap: &callbacks{},
			delete:         &callbacks{},
			deleteAll:      &callbacks{},
			ping:           &callbacks{},
			increment:      &callbacks{},
			decrement:      &callbacks{},
		},
	}
}

func (cw *ClientWrapper) FlushAll() error {
	for _, cb := range cw.registry.flushAll.befores {
		cb.fn(nil, nil)
	}
	err := cw.client.FlushAll()
	for _, cb := range cw.registry.flushAll.afters {
		cb.fn(nil, []any{err})
	}
	return err
}

// Get gets the item for the given key. ErrCacheMiss is returned for a
// memcache cache miss. The key must be at most 250 bytes in length.
func (cw *ClientWrapper) Get(key string) (*memcache.Item, error) {
	for _, cb := range cw.registry.get.befores {
		cb.fn([]any{key}, nil)
	}
	item, err := cw.client.Get(key)
	for _, cb := range cw.registry.get.afters {
		cb.fn([]any{key}, []any{item, err})
	}
	return item, err
}

// Touch updates the expiry for the given key. The seconds parameter is either
// a Unix timestamp or, if seconds is less than 1 month, the number of seconds
// into the future at which time the item will expire. Zero means the item has
// no expiration time. ErrCacheMiss is returned if the key is not in the cache.
// The key must be at most 250 bytes in length.
func (cw *ClientWrapper) Touch(key string, seconds int32) error {
	for _, cb := range cw.registry.touch.befores {
		cb.fn([]any{key, seconds}, nil)
	}
	err := cw.client.Touch(key, seconds)
	for _, cb := range cw.registry.touch.afters {
		cb.fn([]any{key, seconds}, []any{err})
	}
	return err
}

// GetMulti is a batch version of Get. The returned map from keys to
// items may have fewer elements than the input slice, due to memcache
// cache misses. Each key must be at most 250 bytes in length.
// If no error is returned, the returned map will also be non-nil.
func (cw *ClientWrapper) GetMulti(keys []string) (map[string]*memcache.Item, error) {
	for _, cb := range cw.registry.getMulti.befores {
		cb.fn([]any{keys}, nil)
	}
	items, err := cw.client.GetMulti(keys)
	for _, cb := range cw.registry.getMulti.afters {
		cb.fn([]any{keys}, []any{items, err})
	}
	return items, err
}

// Set writes the given item, unconditionally.
func (cw *ClientWrapper) Set(item *memcache.Item) error {
	for _, cb := range cw.registry.set.befores {
		cb.fn([]any{item}, nil)
	}
	err := cw.client.Set(item)
	for _, cb := range cw.registry.set.afters {
		cb.fn([]any{item}, []any{err})
	}
	return err
}

// Add writes the given item, if no value already exists for its
// key. ErrNotStored is returned if that condition is not met.
func (cw *ClientWrapper) Add(item *memcache.Item) error {
	for _, cb := range cw.registry.add.befores {
		cb.fn([]any{item}, nil)
	}
	err := cw.client.Add(item)
	for _, cb := range cw.registry.add.afters {
		cb.fn([]any{item}, []any{err})
	}
	return err
}

// Replace writes the given item, but only if the server *does*
// already hold data for this key
func (cw *ClientWrapper) Replace(item *memcache.Item) error {
	for _, cb := range cw.registry.replace.befores {
		cb.fn([]any{item}, nil)
	}
	err := cw.client.Replace(item)
	for _, cb := range cw.registry.replace.afters {
		cb.fn([]any{item}, []any{err})
	}
	return err
}

// CompareAndSwap writes the given item that was previously returned
// by Get, if the value was neither modified or evicted between the
// Get and the CompareAndSwap calls. The item's Key should not change
// between calls but all other item fields may differ. ErrCASConflict
// is returned if the value was modified in between the
// calls. ErrNotStored is returned if the value was evicted in between
// the calls.
func (cw *ClientWrapper) CompareAndSwap(item *memcache.Item) error {
	for _, cb := range cw.registry.compareAndSwap.befores {
		cb.fn([]any{item}, nil)
	}
	err := cw.client.CompareAndSwap(item)
	for _, cb := range cw.registry.compareAndSwap.afters {
		cb.fn([]any{item}, []any{err})
	}
	return err
}

// Delete deletes the item with the provided key. The error ErrCacheMiss is
// returned if the item didn't already exist in the cache.
func (cw *ClientWrapper) Delete(key string) error {
	for _, cb := range cw.registry.delete.befores {
		cb.fn([]any{key}, nil)
	}
	err := cw.client.Delete(key)
	for _, cb := range cw.registry.delete.afters {
		cb.fn([]any{key}, []any{err})
	}
	return err
}

// DeleteAll deletes all items in the cache.
func (cw *ClientWrapper) DeleteAll() error {
	for _, cb := range cw.registry.deleteAll.befores {
		cb.fn(nil, nil)
	}
	err := cw.client.DeleteAll()
	for _, cb := range cw.registry.deleteAll.afters {
		cb.fn(nil, []any{err})
	}
	return err
}

// Ping checks all instances if they are alive. Returns error if any
// of them is down.
func (cw *ClientWrapper) Ping() error {
	for _, cb := range cw.registry.ping.befores {
		cb.fn(nil, nil)
	}
	err := cw.client.Ping()
	for _, cb := range cw.registry.ping.afters {
		cb.fn(nil, []any{err})
	}
	return err
}

// Increment atomically increments key by delta. The return value is
// the new value after being incremented or an error. If the value
// didn't exist in memcached the error is ErrCacheMiss. The value in
// memcached must be an decimal number, or an error will be returned.
// On 64-bit overflow, the new value wraps around.
func (cw *ClientWrapper) Increment(key string, delta uint64) (uint64, error) {
	for _, cb := range cw.registry.increment.befores {
		cb.fn([]any{key, delta}, nil)
	}
	newValue, err := cw.client.Increment(key, delta)
	for _, cb := range cw.registry.increment.afters {
		cb.fn([]any{key, delta}, []any{newValue, err})
	}
	return newValue, err
}

// Decrement atomically decrements key by delta. The return value is
// the new value after being decremented or an error. If the value
// didn't exist in memcached the error is ErrCacheMiss. The value in
// memcached must be an decimal number, or an error will be returned.
// On underflow, the new value is capped at zero and does not wrap
// around.
func (cw *ClientWrapper) Decrement(key string, delta uint64) (uint64, error) {
	for _, cb := range cw.registry.decrement.befores {
		cb.fn([]any{key, delta}, nil)
	}
	newValue, err := cw.client.Decrement(key, delta)
	for _, cb := range cw.registry.decrement.afters {
		cb.fn([]any{key, delta}, []any{newValue, err})
	}
	return newValue, err
}

// Callback returns callbackRegistry
func (cw *ClientWrapper) Callback() *callbackRegistry {
	return cw.registry
}

type callbackRegistry struct {
	flushAll       *callbacks
	get            *callbacks
	touch          *callbacks
	getMulti       *callbacks
	set            *callbacks
	add            *callbacks
	replace        *callbacks
	compareAndSwap *callbacks
	delete         *callbacks
	deleteAll      *callbacks
	ping           *callbacks
	increment      *callbacks
	decrement      *callbacks
}

func (cr *callbackRegistry) FlushAll() *callbacks {
	return cr.flushAll
}

func (cr *callbackRegistry) Get() *callbacks {
	return cr.get
}

func (cr *callbackRegistry) Touch() *callbacks {
	return cr.touch
}

func (cr *callbackRegistry) GetMulti() *callbacks {
	return cr.getMulti
}

func (cr *callbackRegistry) Set() *callbacks {
	return cr.set
}

func (cr *callbackRegistry) Add() *callbacks {
	return cr.add
}

func (cr *callbackRegistry) Replace() *callbacks {
	return cr.replace
}

func (cr *callbackRegistry) CompareAndSwap() *callbacks {
	return cr.compareAndSwap
}

func (cr *callbackRegistry) Delete() *callbacks {
	return cr.delete
}

func (cr *callbackRegistry) DeleteAll() *callbacks {
	return cr.deleteAll
}

func (cr *callbackRegistry) Ping() *callbacks {
	return cr.ping
}

func (cr *callbackRegistry) Increment() *callbacks {
	return cr.increment
}

func (cr *callbackRegistry) Decrement() *callbacks {
	return cr.decrement
}

type callbacks struct {
	befores handlers
	afters  handlers
}

func (cbs *callbacks) Before() *handlers {
	return &cbs.befores
}

func (cbs *callbacks) After() *handlers {
	return &cbs.afters
}

type handler struct {
	name string
	fn   func(args, results []any)
}

type handlers []*handler

func (hs *handlers) Register(name string, fn func(args, results []any)) {
	*hs = append(*hs, &handler{name, fn})
}

func (hs *handlers) Unregister(name string) {
	for i, h := range *hs {
		if h.name == name {
			*hs = append((*hs)[:i], (*hs)[i+1:]...)
		}
	}
}
