package mdutil

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func SingleValueFromContext(ctx context.Context, key string) string {
	if vals := metadata.ValueFromIncomingContext(ctx, key); len(vals) > 0 {
		return vals[0]
	}
	return ""
}

func MultiValuesFromContext(ctx context.Context, key string) []string {
	return metadata.ValueFromIncomingContext(ctx, key)
}
