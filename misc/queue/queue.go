package queue

import (
	"github.com/eapache/queue/v2"
)

func New[V any]() *queue.Queue[V] {
	return queue.New[V]()
}
