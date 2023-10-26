package i18n

import (
	"context"
)

type funcCtxKey struct{}

func defaultT(messageID string, ops ...LOption) string {
	log("Translation function is not found, use default for messageID: %s", messageID)
	return messageID
}

func NewContext(ctx context.Context, lang string) context.Context {
	if t := GetT(lang); t != nil {
		return context.WithValue(ctx, funcCtxKey{}, t)
	}
	return ctx
}

func FromContext(ctx context.Context) T {
	if t, ok := ctx.Value(funcCtxKey{}).(T); ok {
		return t
	}
	return defaultT
}
