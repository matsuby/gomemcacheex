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
