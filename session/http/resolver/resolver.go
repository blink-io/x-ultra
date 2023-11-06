package resolver

import (
	"context"
	"net/http"
	"time"

	"github.com/blink-io/x/session"
)

type Resolver interface {
	Resolve(m Manager, w http.ResponseWriter, r *http.Request, next http.Handler) error
}

type Manager interface {
	IsRememberMe(context.Context, string) bool
	SetRememberMe(context.Context, string, bool)
	Status(context.Context) session.Status
	Commit(context.Context) (string, time.Time, error)
	Load(context.Context, string) (context.Context, error)
	ErrorFunc(http.ResponseWriter, *http.Request, error)
}
