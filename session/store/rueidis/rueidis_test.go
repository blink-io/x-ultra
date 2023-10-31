package rueidis

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/blink-io/x/redis/rueidis/hooks/debug"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidishook"
	"github.com/stretchr/testify/require"
)

//func init() {
//	os.Setenv("REDIS_TEST_DSN", "redis://localhost:6379/0")
//}

func TestFind(t *testing.T) {
	opt, err := rueidis.ParseURL(os.Getenv("REDIS_TEST_DSN"))
	if err != nil {
		t.Fatal(err)
	}
	client, err := rueidis.NewClient(opt)
	require.NoError(t, err)

	defer client.Close()

	ctx := context.Background()
	r := newRaw(client)

	err = client.Do(ctx, client.B().Flushdb().Build()).Error()
	if err != nil {
		t.Fatal(err)
	}

	setCmd := client.B().Set().Key(r.prefix + "session_token").Value("encoded_data").Ex(99999 * time.Second).Build()
	err = client.Do(ctx, setCmd).Error()
	if err != nil {
		t.Fatal(err)
	}

	b, found, err := r.Find(ctx, "session_token")
	if err != nil {
		t.Fatal(err)
	}
	if found != true {
		t.Fatalf("got %v: expected %v", found, true)
	}
	if bytes.Equal(b, []byte("encoded_data")) == false {
		t.Fatalf("got %v: expected %v", b, []byte("encoded_data"))
	}
}

func TestSaveNew(t *testing.T) {
	opt, err := rueidis.ParseURL(os.Getenv("REDIS_TEST_DSN"))
	if err != nil {
		t.Fatal(err)
	}
	client, err := rueidis.NewClient(opt)
	require.NoError(t, err)

	defer client.Close()

	ctx := context.Background()
	r := newRaw(client)

	err = client.Do(ctx, client.B().Flushdb().Build()).Error()
	if err != nil {
		t.Fatal(err)
	}

	err = r.Commit(ctx, "session_token", []byte("encoded_data"), time.Now().Add(time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	getCmd := client.B().Get().Key(r.prefix + "session_token").Build()
	data, err := client.Do(ctx, getCmd).AsBytes()
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(data, []byte("encoded_data")) == false {
		t.Fatalf("got %v: expected %v", data, []byte("encoded_data"))
	}
}

func TestFindMissing(t *testing.T) {
	opt, err := rueidis.ParseURL(os.Getenv("REDIS_TEST_DSN"))
	if err != nil {
		t.Fatal(err)
	}
	client, err := rueidis.NewClient(opt)
	require.NoError(t, err)

	defer client.Close()

	ctx := context.Background()
	r := New(client)

	err = client.Do(ctx, client.B().Flushdb().Build()).Error()
	if err != nil {
		t.Fatal(err)
	}

	_, found, err := r.Find(ctx, "missing_session_token")
	if err != nil {
		t.Fatalf("got %v: expected %v", err, nil)
	}
	if found != false {
		t.Fatalf("got %v: expected %v", found, false)
	}
}

func TestSaveUpdated(t *testing.T) {
	opt, err := rueidis.ParseURL(os.Getenv("REDIS_TEST_DSN"))
	if err != nil {
		t.Fatal(err)
	}
	client, err := rueidis.NewClient(opt)
	require.NoError(t, err)

	defer client.Close()

	ctx := context.Background()
	r := newRaw(client)

	err = client.Do(ctx, client.B().Flushdb().Build()).Error()
	if err != nil {
		t.Fatal(err)
	}

	setCmd := client.B().Setex().Key(r.prefix + "session_token").Seconds(0).Value("encoded_data").Build()
	err = client.Do(ctx, setCmd).Error()
	if err != nil {
		t.Fatal(err)
	}

	err = r.Commit(ctx, "session_token", []byte("new_encoded_data"), time.Now().Add(time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	getCmd := client.B().Get().Key(r.prefix + "session_token").Build()
	data, err := client.Do(ctx, getCmd).AsBytes()
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(data, []byte("new_encoded_data")) == false {
		t.Fatalf("got %v: expected %v", data, []byte("new_encoded_data"))
	}
}

func TestExpiry(t *testing.T) {
	opt, err := rueidis.ParseURL(os.Getenv("REDIS_TEST_DSN"))
	if err != nil {
		t.Fatal(err)
	}
	client, err := rueidis.NewClient(opt)
	require.NoError(t, err)

	defer client.Close()

	ctx := context.Background()
	r := newRaw(client)

	err = client.Do(ctx, client.B().Flushdb().Build()).Error()
	if err != nil {
		t.Fatal(err)
	}

	err = r.Commit(ctx, "session_token", []byte("encoded_data"), time.Now().Add(100*time.Millisecond))
	if err != nil {
		t.Fatal(err)
	}

	_, found, _ := r.Find(ctx, "session_token")
	if found != true {
		t.Fatalf("got %v: expected %v", found, true)
	}

	time.Sleep(200 * time.Millisecond)
	_, found, _ = r.Find(ctx, "session_token")
	if found != false {
		t.Fatalf("got %v: expected %v", found, false)
	}
}

func TestDelete(t *testing.T) {
	opt, err := rueidis.ParseURL(os.Getenv("REDIS_TEST_DSN"))
	if err != nil {
		t.Fatal(err)
	}
	client, err := rueidis.NewClient(opt)
	require.NoError(t, err)

	client = rueidishook.WithHook(client, debug.New())

	defer client.Close()

	ctx := context.Background()
	r := newRaw(client)

	err = client.Do(ctx, client.B().Flushdb().Build()).Error()
	if err != nil {
		t.Fatal(err)
	}

	setCmd := client.B().Setex().Key(r.prefix + "session_token").Seconds(0).Value("encoded_data").Build()
	err = client.Do(ctx, setCmd).Error()
	if err != nil {
		t.Fatal(err)
	}

	err = r.Delete(ctx, "session_token")
	if err != nil {
		t.Fatal(err)
	}

	getCmd := client.B().Get().Key(r.prefix + "session_token").Build()
	data, err := client.Do(ctx, getCmd).AsBytes()
	if err != rueidis.Nil {
		t.Fatal(err)
	}
	if data != nil {
		t.Fatalf("got %v: expected %v", data, nil)
	}
}

func TestAll(t *testing.T) {
	opt, err := rueidis.ParseURL(os.Getenv("REDIS_TEST_DSN"))
	if err != nil {
		t.Fatal(err)
	}
	client, err := rueidis.NewClient(opt)
	require.NoError(t, err)

	client = rueidishook.WithHook(client, debug.New())

	defer client.Close()

	ctx := context.Background()
	r := newRaw(client)

	err = client.Do(ctx, client.B().Flushdb().Build()).Error()
	if err != nil {
		t.Fatal(err)
	}

	sessions := make(map[string][]byte)
	for i := 0; i < 4; i++ {
		key := fmt.Sprintf("token_%v", i)
		val := []byte(key)
		setCmd := client.B().Set().
			Key(r.prefix + key).
			Value(key).
			Ex(9999 * time.Second).Build()
		err = client.Do(ctx, setCmd).Error()
		if err != nil {
			t.Fatal(err)
		}
		sessions[key] = val
	}

	gotSessions, err := r.All(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for k := range sessions {
		err = r.Delete(ctx, k)
		if err != nil {
			t.Fatal(err)
		}
	}
	if reflect.DeepEqual(sessions, gotSessions) == false {
		t.Fatalf("got %v: expected %v", gotSessions, sessions)
	}
}
