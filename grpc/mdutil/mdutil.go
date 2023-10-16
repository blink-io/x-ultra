package mdutil

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func ValueFromContext(ctx context.Context, key string) string {
	if vals := metadata.ValueFromIncomingContext(ctx, key); len(vals) > 0 {
		return vals[0]
	}
	return ""
}
