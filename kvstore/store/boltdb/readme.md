# Valkeyrie BoltDB

[![GoDoc](https://godoc.org/github.com/kvtools/boltdb?status.png)](https://godoc.org/github.com/kvtools/boltdb)
[![Build Status](https://github.com/kvtools/boltdb/actions/workflows/build.yml/badge.svg)](https://github.com/kvtools/boltdb/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kvtools/boltdb)](https://goreportcard.com/report/github.com/kvtools/boltdb)

[`valkeyrie`](https://github.com/kvtools/valkeyrie) provides a Go native library to store metadata using Distributed Key/Value stores (or common databases).

## Compatibility

A **storage backend** in `valkeyrie` implements (fully or partially) the [Store](https://github.com/kvtools/valkeyrie/blob/master/store/store.go#L69) interface.

| Calls                 | BoltDB |
|-----------------------|:------:|
| Put                   |   🟢   |
| Get                   |   🟢   |
| Delete                |   🟢   |
| Exists                |   🟢   |
| Watch                 |   🔴   |
| WatchTree             |   🔴   |
| NewLock (Lock/Unlock) |   🔴   |
| List                  |   🟢   |
| DeleteTree            |   🟢   |
| AtomicPut             |   🟢   |
| AtomicDelete          |   🟢   |

## Examples

```go
package main

import (
	"context"
	"log"

	"github.com/kvtools/boltdb"
	"github.com/kvtools/valkeyrie"
)

func main() {
	ctx := context.Background()

	config := &boltdb.Config{
		Bucket: "example",
	}

	kv, err := valkeyrie.NewStore(ctx, boltdb.StoreName, []string{"/tmp/mydatabase"}, config)
	if err != nil {
		log.Fatal("Cannot create store")
	}

	key := "foo"

	err = kv.Put(ctx, key, []byte("bar"), nil)
	if err != nil {
		log.Fatalf("Error trying to put value at key: %v", key)
	}

	pair, err := kv.Get(ctx, key, nil)
	if err != nil {
		log.Fatalf("Error trying accessing value at key: %v", key)
	}

	log.Printf("value: %s", string(pair.Value))

	err = kv.Delete(ctx, key)
	if err != nil {
		log.Fatalf("Error trying to delete key %v", key)
	}
}
```
