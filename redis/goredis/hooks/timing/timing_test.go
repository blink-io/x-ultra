package timing

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestHook_Timing_1(t *testing.T) {
	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{"127.0.0.1:6379"},
	})
	c.AddHook(New())

	ctx := context.Background()

	cmds := []redis.Cmder{
		c.Get(ctx, "env"),
		c.ClientInfo(ctx),
		c.ClientID(ctx),
		c.Time(ctx),
		c.ClientList(ctx),
		c.DBSize(ctx),
	}

	for i, cmd := range cmds {
		require.NoError(t, cmd.Err())
		fmt.Printf("res%d: ---> %+v\n", i, cmd)
	}
}

func TestHook_Timing_DoMulti_1(t *testing.T) {
	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{"127.0.0.1:6379"},
	})
	c.AddHook(New())

	ctx := context.Background()

	pp := c.Pipeline()
	pp.Get(ctx, "env")
	pp.ClientInfo(ctx)
	pp.ClientID(ctx)
	pp.Time(ctx)
	pp.ClientList(ctx)
	pp.DBSize(ctx)
	pp.FunctionStats(ctx)

	cmds, err := pp.Exec(ctx)
	require.NoError(t, err)

	for i, cmd := range cmds {
		require.NoError(t, cmd.Err())
		fmt.Printf("res%d: ---> %+v\n", i, cmd)
	}
}
