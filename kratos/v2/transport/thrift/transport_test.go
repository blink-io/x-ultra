package thrift

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/blink-io/x/kratos/v2/transport"
)

func TestTransport_Kind(t *testing.T) {
	o := &Transport{}
	if !reflect.DeepEqual(KindThrift, o.Kind()) {
		t.Errorf("expect %v, got %v", KindThrift, o.Kind())
	}
}

func TestTransport_Endpoint(t *testing.T) {
	v := "hello"
	o := &Transport{endpoint: v}
	if !reflect.DeepEqual(v, o.Endpoint()) {
		t.Errorf("expect %v, got %v", v, o.Endpoint())
	}
}

func TestTransport_Operation(t *testing.T) {
	v := "hello"
	o := &Transport{operation: v}
	if !reflect.DeepEqual(v, o.Operation()) {
		t.Errorf("expect %v, got %v", v, o.Operation())
	}
}

func TestTransport_RequestHeader(t *testing.T) {
	v := headerCarrier{}
	v.Set("a", "1")
	o := &Transport{reqHeader: v}
	if !reflect.DeepEqual("1", o.RequestHeader().Get("a")) {
		t.Errorf("expect %v, got %v", "1", o.RequestHeader().Get("a"))
	}
}

func TestTransport_ReplyHeader(t *testing.T) {
	v := headerCarrier{}
	v.Set("a", "1")
	o := &Transport{replyHeader: v}
	if !reflect.DeepEqual("1", o.ReplyHeader().Get("a")) {
		t.Errorf("expect %v, got %v", "1", o.ReplyHeader().Get("a"))
	}
}

func TestHeaderCarrier_Keys(t *testing.T) {
	v := headerCarrier{}
	v.Set("abb", "1")
	v.Set("bcc", "2")
	want := []string{"Abb", "Bcc"}
	keys := v.Keys()
	sort.Slice(want, func(i, j int) bool {
		return want[i] < want[j]
	})
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	if !reflect.DeepEqual(want, keys) {
		t.Errorf("expect %v, got %v", want, keys)
	}
}

func TestSetOperation(t *testing.T) {
	tr := &Transport{}
	ctx := transport.NewServerContext(context.Background(), tr)
	SetOperation(ctx, "kratos")
	if !reflect.DeepEqual(tr.operation, "kratos") {
		t.Errorf("expect %v, got %v", "kratos", tr.operation)
	}
}
