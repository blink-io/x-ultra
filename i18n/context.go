package i18n

import (
	"context"
)

type fnCtxKey struct{}

func defaultT(messageID string, options ...LOption) string {
	log("Translation function is not found, use default for messageID: %s", messageID)
	return messageID
}

func NewContext(ctx context.Context, lang string) context.Context {
	if t := GetT(lang); t != nil {
		return context.WithValue(ctx, fnCtxKey{}, t)
	}
	return ctx
}

func FromContext(ctx context.Context) T {
	if t, ok := ctx.Value(fnCtxKey{}).(T); ok {
		return t
	}
	return defaultT
}
