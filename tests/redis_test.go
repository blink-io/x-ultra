package tests

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gogo/protobuf/proto"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestRedis_Protobuf_1(t *testing.T) {
	ctx := context.Background()

	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{":6379"},
	})

	pb := durationpb.New(time.Second * 100)
	pbData, err := proto.Marshal(pb)
	require.NoError(t, err)

	hstr := hex.EncodeToString(pbData)
	fmt.Println("Hex: ", hstr)
	fmt.Println("xxx: ", fmt.Sprint(pbData))

	errv := c.Set(ctx, "dur_pb", pb, 0).Err()
	require.NoError(t, errv)
}

func TestRedisServer_Mini(t *testing.T) {
	s := miniredis.RunT(t)

	ctx := context.Background()

	// Optionally set some keys your code expects:
	s.Set("foo", "bar")
	s.HSet("some", "other", "key")

	// Run your code and see if it behaves.
	// An example using the redigo library from "github.com/gomodule/redigo/redis":
	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{":6379"},
	})
	c.Set(ctx, "foo", "bar", 0)

	// Optionally check values in redis...
	if got, err := s.Get("foo"); err != nil || got != "bar" {
		t.Error("'foo' has the wrong value")
	}
	// ... or use a helper for that:
	s.CheckGet(t, "foo", "bar")

	// TTL and expiration:
	s.Set("foo", "bar")
	s.SetTTL("foo", 10*time.Second)
	s.FastForward(11 * time.Second)
	if s.Exists("foo") {

	}

}
