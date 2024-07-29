package bun

import (
	"context"
)

func Struct[M ModelType](ctx context.Context, db RawIDB, ops ...DoSelectOption) (ModelSlice[M], error) {
	var ms ModelSlice[M]
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectAll, q, o)
	err := q.Scan(ctx, &ms)
	return ms, err
}

func StructSQL[M ModelType](ctx context.Context, db RawIDB, query string, args ...any) (ModelSlice[M], error) {
	var ms ModelSlice[M]
	err := db.NewRaw(query, args...).Scan(ctx, &ms)
	return ms, err
}
