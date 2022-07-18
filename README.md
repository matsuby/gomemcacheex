# gomemcacheex

[![Go Reference](https://pkg.go.dev/badge/github.com/matsuby/gomemcacheex.svg)](https://pkg.go.dev/github.com/matsuby/gomemcacheex)

## About
gomemcacheex povides wrapper for [Go memcached client (bradfitz/gomemcache)](https://github.com/bradfitz/gomemcache) that name `ClientWrapper`.

`ClientWrapper` can register callback functions for each methods(`Get`, `Set`, `Delete`, etc) similar to [GORM](https://gorm.io/docs/write_plugins.html).

## Installing
```
go get github.com/matsuby/gomemcacheex
```

## Usage
```
package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/matsuby/gomemcacheex/memcacheex"
)

func main() {
	// create wrapper for memcached client
	cw := memcacheex.NewClientWrapper(memcache.New("localhost:11211"))

	// register callback functions
	cw.Callback().Set().Before().Register("gomemcacheex:set-before", func(args, results []any) {
		fmt.Println("--- Set: Before")
		fmt.Println(args...)
	})
	cw.Callback().Get().After().Register("gomemcacheex:get-after", func(args, results []any) {
		fmt.Println("--- Get: After")
		fmt.Println(args...)
		fmt.Println(results...)
	})

	// call methods, then invoke registered callback functions
	_ = cw.Set(&memcache.Item{Key: "test_key", Value: []byte("test_value")})
	_, _ = cw.Get("test_key")
}
```

## Bonus
gomemcacheex provodes [interface](memcacheex/interface.go) and [mock](memcacheex/mock.go) for memcached client. These may help you with your unit tests.