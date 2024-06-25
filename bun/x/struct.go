package x

import (
	"context"

	rdb "github.com/blink-io/x/bun"
)

func Struct[M ModelType](ctx context.Context, db rdb.RawIDB, ops ...SelectOption) (ModelSlice[M], error) {
	var ms ModelSlice[M]
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)
	err := q.Scan(ctx, &ms)
	return ms, err
}

func StructSQL[M ModelType](ctx context.Context, db rdb.RawIDB, query string, args ...any) (ModelSlice[M], error) {
	var ms ModelSlice[M]
	err := db.NewRaw(query, args...).Scan(ctx, &ms)
	return ms, err
}
