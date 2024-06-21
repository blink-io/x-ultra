package x

import (
	"context"

	rdb "github.com/blink-io/x/bun"
)

type ValueType interface {
	string |
		~int | ~uint |
		~int8 | ~uint8 |
		~int16 | ~uint16 |
		~int32 | ~uint32 |
		~int64 | ~uint64 |
		~float32 | ~float64 |
		~bool
}

type TypeSlice[T ValueType] []T

func Type[T ValueType](ctx context.Context, db rdb.RawIDB, table string, column string, ops ...SelectOption) (TypeSlice[T], error) {
	var ts TypeSlice[T]
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)
	err := q.Column(column).
		Table(table).
		Scan(ctx, &ts)
	return ts, err
}
