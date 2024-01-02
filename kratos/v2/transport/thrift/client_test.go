package thrift

import (
	"context"
	"testing"

	i18nthrift "github.com/blink-io/x/i18n/thrift"
)

func TestClient(t *testing.T) {
	ctx := context.Background()
	conn, err := Dial(
		WithEndpoint("localhost:7700"),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := i18nthrift.NewI18NClient(conn.Client)

	req := &i18nthrift.ListLanguagesRequest{
		Languages: []string{"zh-Hans"},
	}
	reply, err := client.ListLanguages(ctx, req)
	//t.Log(err)
	if err != nil {
		t.Errorf("failed to call: %v", err)
	}
	t.Log(reply.Timestamp, reply.Entries)
}
