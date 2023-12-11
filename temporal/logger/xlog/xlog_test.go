package xlog

import (
	"testing"

	xlog "github.com/blink-io/x/log"
)

func TestZap_1(t *testing.T) {
	msgs := map[string][]any{
		"msg1": {"ver"},
		"msg2": {"ver", "v1"},
		"msg3": {"ver", "v1", "score"},
		"msg4": {"ver", "v1", "score", 99.0, "kkk", true},
		"msg5": {"ver", "v1", "score", 99.0, "kkk", true, 999},
		"msg6": {"ver", "v1", "score", 99.0, "kkk", true, 999, "ok"},
		//"msg7": {"ver", "v1", "score", 99.0, "kkk"},
		//"msg8": {"ver", "v1", "score", 99.0, "kkk"},
		//"msg9": {"ver", "v1", "score", 99.0, "kkk"},
	}
	log := NewLogger(xlog.DefaultLogger)

	for k, v := range msgs {
		log.Info(k, v...)
	}
}
