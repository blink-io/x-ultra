package debug

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidishook"
	"github.com/stretchr/testify/require"
)

func TestHook_Timing_1(t *testing.T) {
	c, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
	})
	require.NoError(t, err)

	c = rueidishook.WithHook(c, New())

	cmds := []rueidis.Completed{
		c.B().Get().Key("env").Build(),
		c.B().ClientSetname().ConnectionName("test-from-linux-ser6v").Build(),
		c.B().ClientGetname().Build(),
		c.B().ClientId().Build(),
	}

	ctx := context.Background()

	for i, cmd := range cmds {
		res, err := c.Do(ctx, cmd).ToAny()
		require.NoError(t, err)
		fmt.Printf("res%d: ---> %+v\n", i, res)
	}
}

func TestHook_Timing_DoMulti_1(t *testing.T) {
	c, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
	})
	require.NoError(t, err)

	c = rueidishook.WithHook(c, New())

	cmds := []rueidis.Completed{
		c.B().Get().Key("env").Build(),
		c.B().ClientSetname().ConnectionName("test-from-linux-ser6v").Build(),
		c.B().ClientGetname().Build(),
		c.B().ClientId().Build(),
	}

	ctx := context.Background()

	ress := c.DoMulti(ctx, cmds...)

	for i, res := range ress {
		val, err := res.ToAny()
		require.NoError(t, err)
		fmt.Printf("res%d: ---> %+v\n", i, val)
	}
}
