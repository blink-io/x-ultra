package encoding

import (
	"sync"
)

var registers = make(map[string]Codec)

var mu sync.Mutex

func Register(name string, codec Codec) {
	mu.Lock()
	defer mu.Unlock()

	if codec == nil {
		panic("session encoding: codec can not be nil")
	}

	if _, dup := registers[name]; dup {
		panic("session encoding: register called twice for codec " + name)
	}

	registers[name] = codec
}

func Get(name string) (Codec, bool) {
	mu.Lock()
	defer mu.Unlock()

	c, ok := registers[name]
	return c, ok
}
