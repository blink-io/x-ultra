package g

import (
	"time"

	"github.com/nsqio/go-nsq"
)

type (
	Consumer interface {
		AddHandler(handler nsq.Handler)

		AddConcurrentHandlers(handler nsq.Handler, concurrency int)
	}

	Producer interface {
		Ping() error

		DeferredPublish(topic string, delay time.Duration, body []byte) error

		DeferredPublishAsync(topic string, delay time.Duration, body []byte,
			doneChan chan *nsq.ProducerTransaction, args ...interface{}) error

		MultiPublish(topic string, body [][]byte) error

		MultiPublishAsync(topic string, body [][]byte, doneChan chan *nsq.ProducerTransaction,
			args ...interface{}) error

		Publish(topic string, body []byte) error

		PublishAsync(topic string, body []byte, doneChan chan *nsq.ProducerTransaction,
			args ...interface{}) error
	}
)

var _ Consumer = (*nsq.Consumer)(nil)
var _ Producer = (*nsq.Producer)(nil)
