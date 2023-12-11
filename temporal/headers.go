package temporal

import (
	"context"
)

type Headers map[string]string

func (h Headers) GetHeaders(ctx context.Context) (map[string]string, error) {
	return h, nil
}
