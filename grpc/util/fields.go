package util

import (
	"context"
)

type fieldsCtxMarker struct{}

var (
	// fieldsCtxMarkerKey is the Context value marker that is used by logging middleware to read and write logging fields into context.
	fieldsCtxMarkerKey = &fieldsCtxMarker{}
)

// Fields loosely represents key value pairs that adds context to log lines. The key has to be type of string, whereas
// value can be an arbitrary object.
type Fields []any

// Iterator returns iterator that allows iterating over pair of elements representing field.
// If number of elements is uneven, last element won't be included will be assumed as key with empty string value.
// If key is not string, At will panic.
func (f Fields) Iterator() *iter {
	// We start from -2 as we iterate over two items per iteration and first iteration will advance iterator to 0.
	return &iter{i: -2, f: f}
}

type iter struct {
	f Fields
	i int
}

func (i *iter) Next() bool {
	if i.i >= len(i.f) {
		return false
	}

	i.i += 2
	return i.i < len(i.f)
}

func (i *iter) At() (k string, v any) {
	if i.i < 0 || i.i >= len(i.f) {
		return "", ""
	}

	if i.i+1 == len(i.f) {
		// Non even number of elements, add empty string.
		return i.f[i.i].(string), ""
	}
	return i.f[i.i].(string), i.f[i.i+1]
}

func (f *Fields) Delete(key string) {
	i := f.Iterator()
	for i.Next() {
		k, _ := i.At()
		if k != key {
			continue
		}
		*f = append((*f)[:i.i], (*f)[i.i+2:]...)
		return
	}
}

// WithUnique returns copy of fields which is the union of all unique keys.
// Any duplicates in the added or current fields will be deduplicated where first occurrence takes precedence.
func (f Fields) WithUnique(add Fields) Fields {
	if len(add) == 0 {
		n := make(Fields, len(f))
		copy(n, f)
		return n
	}

	existing := map[any]struct{}{}
	i := f.Iterator()
	for i.Next() {
		k, _ := i.At()
		existing[k] = struct{}{}
	}

	n := make(Fields, len(f), len(f)+len(add))
	copy(n, f)

	a := add.Iterator()
	for a.Next() {
		k, v := a.At()
		if _, ok := existing[k]; ok {
			continue
		}
		n = append(n, k, v)
	}
	return n
}

// AppendUnique appends (can reuse array!) fields which does not occur in existing fields slice.
func (f Fields) AppendUnique(add Fields) Fields {
	if len(add) == 0 {
		return f
	}

	a := add.Iterator()
NextAddField:
	for a.Next() {
		k, v := a.At()
		i := f.Iterator()
		for i.Next() {
			fk, _ := i.At()
			if fk == k {
				continue NextAddField
			}
		}
		f = append(f, k, v)
	}
	return f
}

// ExtractFields returns logging fields from the context.
// Fields can be added from the context using InjectFields. For example logging interceptor adds some (base) fields
// into context when used.
// If there are no fields in the context, it returns an empty Fields value.
// Extracted fields are useful to construct your own logger that has fields from gRPC interceptors.
func ExtractFields(ctx context.Context) Fields {
	t, ok := ctx.Value(fieldsCtxMarkerKey).(Fields)
	if !ok {
		return nil
	}
	n := make(Fields, len(t))
	copy(n, t)
	return n
}

// InjectFields allows adding fields to any existing Fields that will be used by the logging interceptor or can be
// extracted further in ExtractFields.
// For explicitness, in case of duplicates, the newest field occurrence wins. This allows nested components to update
// popular fields like grpc.component (e.g. server invoking gRPC client).
//
// Don't overuse mutation of fields to avoid surprises.
func InjectFields(ctx context.Context, f Fields) context.Context {
	return context.WithValue(ctx, fieldsCtxMarkerKey, f.WithUnique(ExtractFields(ctx)))
}
