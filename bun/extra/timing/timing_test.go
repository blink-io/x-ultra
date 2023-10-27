package timing

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/uptrace/bun"
)

func TestQueryHook_1(t *testing.T) {
	ctx := context.Background()
	event := &bun.QueryEvent{}
	h := &hook{}
	ctx = h.BeforeQuery(ctx, event)
	time.Sleep(3 * time.Second)
	h.AfterQuery(ctx, event)

	fmt.Println("done")
}
