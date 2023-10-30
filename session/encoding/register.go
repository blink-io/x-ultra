package encoding

import (
	"sync"
)

var registers = make(map[string]Codec)

var mu sync.Mutex

func Register(name string, codec Codec) {
	mu.Lock()
	defer mu.Unlock()

	registers[name] = codec
}
