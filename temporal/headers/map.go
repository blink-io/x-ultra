package headers

import (
	"context"
)

type Map map[string]string

func (m Map) GetHeaders(ctx context.Context) (map[string]string, error) {
	return m, nil
}
