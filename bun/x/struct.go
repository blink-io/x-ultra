package x

import (
	"context"

	rdb "github.com/blink-io/x/bun"
)

//type ValueType interface {
//	any | ~string | ~bool |
//		~int | ~uint |
//		~int8 | ~uint8 |
//		~int16 | ~uint16 |
//		~int32 | ~uint32 |
//		~int64 | ~uint64 |
//		~float32 | ~float64 |
//		~*string | ~*bool |
//		~*int | ~*uint |
//		~*int8 | ~*uint8 |
//		~*int16 | ~*uint16 |
//		~*int32 | ~*uint32 |
//		~*int64 | ~*uint64 |
//		~*float32 | ~*float64 |
//		sql.Null[string] | sql.Null[bool] |
//		sql.Null[int] | sql.Null[uint] |
//		sql.Null[int8] | sql.Null[uint8] |
//		sql.Null[int16] | sql.Null[uint16] |
//		sql.Null[int32] | sql.Null[uint32] |
//		sql.Null[int64] | sql.Null[uint64] |
//		sql.Null[float32] | sql.Null[float64]
//}

type ValueType = any

type TypeSlice[T ValueType] []T

type ColumnValue[V ValueType] struct {
	Column string
	Value  V
}

func NewColumnValue[V ValueType](column string, value V) *ColumnValue[V] {
	return &ColumnValue[V]{
		Column: column,
		Value:  value,
	}
}

type Tuple2[T1 ValueType, T2 ValueType] struct {
	T1 T1
	T2 T2
}

type Tuple3[T1 ValueType, T2 ValueType, T3 ValueType] struct {
	T1 T1
	T2 T2
	T3 T3
}

type Tuple4[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType] struct {
	T1 T1
	T2 T2
	T3 T3
	T4 T4
}

type Tuple5[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType] struct {
	T1 T1
	T2 T2
	T3 T3
	T4 T4
	T5 T5
}

type Tuple6[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType] struct {
	T1 T1
	T2 T2
	T3 T3
	T4 T4
	T5 T5
	T6 T6
}

type Tuple7[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType] struct {
	T1 T1
	T2 T2
	T3 T3
	T4 T4
	T5 T5
	T6 T6
	T7 T7
}

type Tuple8[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType,
	T6 ValueType, T7 ValueType, T8 ValueType] struct {
	T1 T1
	T2 T2
	T3 T3
	T4 T4
	T5 T5
	T6 T6
	T7 T7
	T8 T8
}

type Tuple9[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType,
	T6 ValueType, T7 ValueType, T8 ValueType, T9 ValueType] struct {
	T1 T1
	T2 T2
	T3 T3
	T4 T4
	T5 T5
	T6 T6
	T7 T7
	T8 T8
	T9 T9
}

func Type[T ValueType](ctx context.Context, db rdb.RawIDB,
	table string, column string, ops ...SelectOption) (TypeSlice[T], error) {
	var ts TypeSlice[T]
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)
	err := q.Column(column).
		Table(table).
		Scan(ctx, &ts)
	return ts, err
}

func TypeTuple2[T1 ValueType, T2 ValueType](ctx context.Context, db rdb.RawIDB,
	table string, col1, col2 string, ops ...SelectOption) ([]*Tuple2[T1, T2], error) {
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)

	var v1s []T1
	var v2s []T2
	err := q.Column(col1, col2).
		Table(table).
		Scan(ctx, &v1s, &v2s)

	maxLen := max(len(v1s), len(v2s))
	ts := make([]*Tuple2[T1, T2], maxLen)
	for i := 0; i < maxLen; i++ {
		ts[i] = &Tuple2[T1, T2]{}
	}
	for i, v1 := range v1s {
		ti := ts[i]
		ti.T1 = v1
	}
	for i, v2 := range v2s {
		ti := ts[i]
		ti.T2 = v2
	}
	return ts, err
}

func TypeTuple3[T1 ValueType, T2 ValueType, T3 ValueType](ctx context.Context, db rdb.RawIDB,
	table string, col1, col2, col3 string, ops ...SelectOption) ([]*Tuple3[T1, T2, T3], error) {
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	err := q.Column(col1, col2, col3).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s)

	maxLen := max(len(v1s), len(v2s), len(v3s))
	ts := make([]*Tuple3[T1, T2, T3], maxLen)
	for i := 0; i < maxLen; i++ {
		ts[i] = &Tuple3[T1, T2, T3]{}
	}
	for i, v1 := range v1s {
		ti := ts[i]
		ti.T1 = v1
	}
	for i, v2 := range v2s {
		ti := ts[i]
		ti.T2 = v2
	}
	for i, v3 := range v3s {
		ti := ts[i]
		ti.T3 = v3
	}
	return ts, err
}

func TypeTuple4[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType](ctx context.Context, db rdb.RawIDB,
	table string, col1, col2, col3, col4 string, ops ...SelectOption) ([]*Tuple4[T1, T2, T3, T4], error) {
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	err := q.Column(col1, col2, col3, col4).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s)

	maxLen := max(len(v1s), len(v2s), len(v3s), len(v4s))
	ts := make([]*Tuple4[T1, T2, T3, T4], maxLen)
	for i := 0; i < maxLen; i++ {
		ts[i] = &Tuple4[T1, T2, T3, T4]{}
	}
	for i, v1 := range v1s {
		ti := ts[i]
		ti.T1 = v1
	}
	for i, v2 := range v2s {
		ti := ts[i]
		ti.T2 = v2
	}
	for i, v3 := range v3s {
		ti := ts[i]
		ti.T3 = v3
	}
	for i, v4 := range v4s {
		ti := ts[i]
		ti.T4 = v4
	}
	return ts, err
}

func TypeTuple5[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType](
	ctx context.Context, db rdb.RawIDB,
	table string, col1, col2, col3, col4, col5 string, ops ...SelectOption) ([]*Tuple5[T1, T2, T3, T4, T5], error) {
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	err := q.Column(col1, col2, col3, col4, col5).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s)

	maxLen := max(len(v1s), len(v2s), len(v3s), len(v4s), len(v5s))
	ts := make([]*Tuple5[T1, T2, T3, T4, T5], maxLen)
	for i := 0; i < maxLen; i++ {
		ts[i] = &Tuple5[T1, T2, T3, T4, T5]{}
	}
	for i, v1 := range v1s {
		ti := ts[i]
		ti.T1 = v1
	}
	for i, v2 := range v2s {
		ti := ts[i]
		ti.T2 = v2
	}
	for i, v3 := range v3s {
		ti := ts[i]
		ti.T3 = v3
	}
	for i, v4 := range v4s {
		ti := ts[i]
		ti.T4 = v4
	}
	for i, v5 := range v5s {
		ti := ts[i]
		ti.T5 = v5
	}
	return ts, err
}

func TypeTuple6[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType](
	ctx context.Context, db rdb.RawIDB,
	table string, col1, col2, col3, col4, col5, col6 string,
	ops ...SelectOption) ([]*Tuple6[T1, T2, T3, T4, T5, T6], error) {
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	var v6s []T6
	err := q.Column(col1, col2, col3, col4, col5, col6).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s, &v6s)

	maxLen := max(len(v1s), len(v2s), len(v3s), len(v4s), len(v5s), len(v6s))
	ts := make([]*Tuple6[T1, T2, T3, T4, T5, T6], maxLen)
	for i := 0; i < maxLen; i++ {
		ts[i] = &Tuple6[T1, T2, T3, T4, T5, T6]{}
	}
	for i, v1 := range v1s {
		ti := ts[i]
		ti.T1 = v1
	}
	for i, v2 := range v2s {
		ti := ts[i]
		ti.T2 = v2
	}
	for i, v3 := range v3s {
		ti := ts[i]
		ti.T3 = v3
	}
	for i, v4 := range v4s {
		ti := ts[i]
		ti.T4 = v4
	}
	for i, v5 := range v5s {
		ti := ts[i]
		ti.T5 = v5
	}
	for i, v6 := range v6s {
		ti := ts[i]
		ti.T6 = v6
	}
	return ts, err
}

func TypeTuple7[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType](
	ctx context.Context, db rdb.RawIDB,
	table string, col1, col2, col3, col4, col5, col6, col7 string,
	ops ...SelectOption) ([]*Tuple7[T1, T2, T3, T4, T5, T6, T7], error) {
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	var v6s []T6
	var v7s []T7
	err := q.Column(col1, col2, col3, col4, col5, col6, col7).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s, &v6s, &v7s)

	maxLen := max(len(v1s), len(v2s), len(v3s), len(v4s), len(v5s), len(v6s), len(v7s))
	ts := make([]*Tuple7[T1, T2, T3, T4, T5, T6, T7], maxLen)
	for i := 0; i < maxLen; i++ {
		ts[i] = &Tuple7[T1, T2, T3, T4, T5, T6, T7]{}
	}
	for i, v1 := range v1s {
		ti := ts[i]
		ti.T1 = v1
	}
	for i, v2 := range v2s {
		ti := ts[i]
		ti.T2 = v2
	}
	for i, v3 := range v3s {
		ti := ts[i]
		ti.T3 = v3
	}
	for i, v4 := range v4s {
		ti := ts[i]
		ti.T4 = v4
	}
	for i, v5 := range v5s {
		ti := ts[i]
		ti.T5 = v5
	}
	for i, v6 := range v6s {
		ti := ts[i]
		ti.T6 = v6
	}

	for i, v7 := range v7s {
		ti := ts[i]
		ti.T7 = v7
	}
	return ts, err
}

func TypeTuple8[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType, T8 ValueType](
	ctx context.Context, db rdb.RawIDB,
	table string, col1, col2, col3, col4, col5, col6, col7, col8 string,
	ops ...SelectOption) ([]*Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], error) {
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	var v6s []T6
	var v7s []T7
	var v8s []T8
	err := q.Column(col1, col2, col3, col4, col5, col6, col7, col8).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s, &v6s, &v7s, &v8s)

	maxLen := max(len(v1s), len(v2s), len(v3s), len(v4s), len(v5s), len(v6s), len(v7s), len(v8s))
	ts := make([]*Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], maxLen)
	for i := 0; i < maxLen; i++ {
		ts[i] = &Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{}
	}
	for i, v1 := range v1s {
		ti := ts[i]
		ti.T1 = v1
	}
	for i, v2 := range v2s {
		ti := ts[i]
		ti.T2 = v2
	}
	for i, v3 := range v3s {
		ti := ts[i]
		ti.T3 = v3
	}
	for i, v4 := range v4s {
		ti := ts[i]
		ti.T4 = v4
	}
	for i, v5 := range v5s {
		ti := ts[i]
		ti.T5 = v5
	}
	for i, v6 := range v6s {
		ti := ts[i]
		ti.T6 = v6
	}
	for i, v7 := range v7s {
		ti := ts[i]
		ti.T7 = v7
	}
	for i, v8 := range v8s {
		ti := ts[i]
		ti.T8 = v8
	}
	return ts, err
}

func TypeTuple9[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType, T8 ValueType, T9 ValueType](
	ctx context.Context, db rdb.RawIDB,
	table string, col1, col2, col3, col4, col5, col6, col7, col8, col9 string,
	ops ...SelectOption) ([]*Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], error) {
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	var v6s []T6
	var v7s []T7
	var v8s []T8
	var v9s []T9
	err := q.Column(col1, col2, col3, col4, col5, col6, col7, col8, col9).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s, &v6s, &v7s, &v8s, &v9s)

	maxLen := max(len(v1s), len(v2s), len(v3s), len(v4s), len(v5s), len(v6s), len(v7s), len(v8s), len(v9s))
	ts := make([]*Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], maxLen)
	for i := 0; i < maxLen; i++ {
		ts[i] = &Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{}
	}
	for i, v1 := range v1s {
		ti := ts[i]
		ti.T1 = v1
	}
	for i, v2 := range v2s {
		ti := ts[i]
		ti.T2 = v2
	}
	for i, v3 := range v3s {
		ti := ts[i]
		ti.T3 = v3
	}
	for i, v4 := range v4s {
		ti := ts[i]
		ti.T4 = v4
	}
	for i, v5 := range v5s {
		ti := ts[i]
		ti.T5 = v5
	}
	for i, v6 := range v6s {
		ti := ts[i]
		ti.T6 = v6
	}
	for i, v7 := range v7s {
		ti := ts[i]
		ti.T7 = v7
	}
	for i, v8 := range v8s {
		ti := ts[i]
		ti.T8 = v8
	}
	for i, v9 := range v9s {
		ti := ts[i]
		ti.T9 = v9
	}
	return ts, err
}
