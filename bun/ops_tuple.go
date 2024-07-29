package bun

import (
	"context"
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

func Type[T ValueType](ctx context.Context, db RawIDB,
	table string, column string, ops ...DoSelectOption) (TypeSlice[T], error) {
	var ts TypeSlice[T]
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectAll, q, o)
	afterSelectOptions(q, table, column)
	err := q.Scan(ctx, &ts)
	return ts, err
}

func TypeTuple2[T1 ValueType, T2 ValueType](ctx context.Context, db RawIDB,
	table string, col1, col2 string, ops ...DoSelectOption) ([]*Tuple2[T1, T2], error) {
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectAll, q, o)
	afterSelectOptions(q, table, col1, col2)

	var v1s []T1
	var v2s []T2
	err := q.Scan(ctx, &v1s, &v2s)
	if err != nil {
		return nil, err
	}
	return handleTuple2Values(v1s, v2s), nil
}

func TypeTuple2SQL[T1 ValueType, T2 ValueType](ctx context.Context, db RawIDB, query string, args ...any) ([]*Tuple2[T1, T2], error) {
	var v1s []T1
	var v2s []T2
	err := db.NewRaw(query, args...).Scan(ctx, &v1s, &v2s)
	if err != nil {
		return nil, err
	}
	return handleTuple2Values(v1s, v2s), nil
}

func handleTuple2Values[T1 ValueType, T2 ValueType](v1s []T1, v2s []T2) []*Tuple2[T1, T2] {
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
	return ts
}

func TypeTuple3[T1 ValueType, T2 ValueType, T3 ValueType](ctx context.Context, db RawIDB,
	table string, col1, col2, col3 string, ops ...DoSelectOption) ([]*Tuple3[T1, T2, T3], error) {
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectAll, q, o)
	afterSelectOptions(q, table, col1, col2, col3)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	if err := q.Scan(ctx, &v1s, &v2s, &v3s); err != nil {
		return nil, err
	}
	return handleTuple3Values(v1s, v2s, v3s), nil
}

func TypeTuple3SQL[T1 ValueType, T2 ValueType, T3 ValueType](ctx context.Context, db RawIDB, query string, args ...any) ([]*Tuple3[T1, T2, T3], error) {
	var v1s []T1
	var v2s []T2
	var v3s []T3
	if err := db.NewRaw(query, args...).
		Scan(ctx, &v1s, &v2s, &v3s); err != nil {
		return nil, err
	}
	return handleTuple3Values(v1s, v2s, v3s), nil
}

func handleTuple3Values[T1 ValueType, T2 ValueType, T3 ValueType](v1s []T1, v2s []T2, v3s []T3) []*Tuple3[T1, T2, T3] {
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
	return ts
}

func TypeTuple4[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType](ctx context.Context, db RawIDB,
	table string, col1, col2, col3, col4 string, ops ...DoSelectOption) ([]*Tuple4[T1, T2, T3, T4], error) {
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectAll, q, o)
	afterSelectOptions(q, table, col1, col2, col3, col4)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	if err := q.Column(col1, col2, col3, col4).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s); err != nil {
		return nil, err
	}
	return handleTuple4Values(v1s, v2s, v3s, v4s), nil
}

func TypeTuple4SQL[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType](
	ctx context.Context, db RawIDB, query string, args ...any) ([]*Tuple4[T1, T2, T3, T4], error) {
	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	if err := db.NewRaw(query, args...).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s); err != nil {
		return nil, err
	}
	return handleTuple4Values(v1s, v2s, v3s, v4s), nil
}

func handleTuple4Values[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType](
	v1s []T1, v2s []T2, v3s []T3, v4s []T4) []*Tuple4[T1, T2, T3, T4] {
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
	return ts
}

func TypeTuple5[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType](
	ctx context.Context, db RawIDB,
	table string, col1, col2, col3, col4, col5 string, ops ...DoSelectOption) ([]*Tuple5[T1, T2, T3, T4, T5], error) {
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectAll, q, o)
	afterSelectOptions(q, table, col1, col2, col3, col4, col5)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	if err := q.Column(col1, col2, col3, col4, col5).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s); err != nil {
		return nil, err
	}
	return handleTuple5Values(v1s, v2s, v3s, v4s, v5s), nil
}

func TypeTuple5SQL[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType](
	ctx context.Context, db RawIDB, query string, args ...any) ([]*Tuple5[T1, T2, T3, T4, T5], error) {
	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	if err := db.NewRaw(query, args...).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s); err != nil {
		return nil, err
	}
	return handleTuple5Values(v1s, v2s, v3s, v4s, v5s), nil
}

func handleTuple5Values[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType](
	v1s []T1, v2s []T2, v3s []T3, v4s []T4, v5s []T5) []*Tuple5[T1, T2, T3, T4, T5] {
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
	return ts
}

func TypeTuple6[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType](
	ctx context.Context, db RawIDB,
	table string, col1, col2, col3, col4, col5, col6 string,
	ops ...DoSelectOption) ([]*Tuple6[T1, T2, T3, T4, T5, T6], error) {
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectAll, q, o)
	afterSelectOptions(q, table, col1, col2, col3, col4, col5, col6)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	var v6s []T6
	if err := q.Column(col1, col2, col3, col4, col5, col6).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s, &v6s); err != nil {
		return nil, err
	}
	return handleTuple6Values(v1s, v2s, v3s, v4s, v5s, v6s), nil
}

func TypeTuple6SQL[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType](
	ctx context.Context, db RawIDB, query string, args ...any) ([]*Tuple6[T1, T2, T3, T4, T5, T6], error) {
	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	var v6s []T6
	if err := db.NewRaw(query, args...).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s, &v6s); err != nil {
		return nil, err
	}
	return handleTuple6Values(v1s, v2s, v3s, v4s, v5s, v6s), nil
}

func handleTuple6Values[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType](
	v1s []T1, v2s []T2, v3s []T3, v4s []T4, v5s []T5, v6s []T6) []*Tuple6[T1, T2, T3, T4, T5, T6] {
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
	return ts
}

func TypeTuple7[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType](
	ctx context.Context, db RawIDB,
	table string, col1, col2, col3, col4, col5, col6, col7 string,
	ops ...DoSelectOption) ([]*Tuple7[T1, T2, T3, T4, T5, T6, T7], error) {
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectAll, q, o)
	afterSelectOptions(q, table, col1, col2, col3, col4, col5, col6, col7)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	var v6s []T6
	var v7s []T7
	if err := q.Column(col1, col2, col3, col4, col5, col6, col7).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s, &v6s, &v7s); err != nil {
		return nil, err
	}
	return handleTuple7Values(v1s, v2s, v3s, v4s, v5s, v6s, v7s), nil
}

func TypeTuple7SQL[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType](
	ctx context.Context, db RawIDB, query string, args ...any) ([]*Tuple7[T1, T2, T3, T4, T5, T6, T7], error) {
	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	var v6s []T6
	var v7s []T7
	if err := db.NewRaw(query, args...).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s, &v6s, &v7s); err != nil {
		return nil, err
	}
	return handleTuple7Values(v1s, v2s, v3s, v4s, v5s, v6s, v7s), nil
}

func handleTuple7Values[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType](
	v1s []T1, v2s []T2, v3s []T3, v4s []T4, v5s []T5, v6s []T6, v7s []T7) []*Tuple7[T1, T2, T3, T4, T5, T6, T7] {
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
	return ts
}

func TypeTuple8[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType, T8 ValueType](
	ctx context.Context, db RawIDB,
	table string, col1, col2, col3, col4, col5, col6, col7, col8 string,
	ops ...DoSelectOption) ([]*Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], error) {
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectAll, q, o)
	afterSelectOptions(q, table, col1, col2, col3, col4, col5, col6, col7, col8)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	var v6s []T6
	var v7s []T7
	var v8s []T8
	if err := q.Column(col1, col2, col3, col4, col5, col6, col7, col8).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s, &v6s, &v7s, &v8s); err != nil {
		return nil, err
	}
	return handleTuple8Values(v1s, v2s, v3s, v4s, v5s, v6s, v7s, v8s), nil
}

func TypeTuple8SQL[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType, T8 ValueType](
	ctx context.Context, db RawIDB,
	query string, args ...any) ([]*Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], error) {
	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	var v6s []T6
	var v7s []T7
	var v8s []T8
	if err := db.NewRaw(query, args...).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s, &v6s, &v7s, &v8s); err != nil {
		return nil, err
	}
	return handleTuple8Values(v1s, v2s, v3s, v4s, v5s, v6s, v7s, v8s), nil
}

func handleTuple8Values[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType, T8 ValueType](
	v1s []T1, v2s []T2, v3s []T3, v4s []T4, v5s []T5, v6s []T6, v7s []T7, v8s []T8) []*Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
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
	return ts
}

func TypeTuple9[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType, T8 ValueType, T9 ValueType](
	ctx context.Context, db RawIDB,
	table string, col1, col2, col3, col4, col5, col6, col7, col8, col9 string,
	ops ...DoSelectOption) ([]*Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], error) {
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectAll, q, o)
	afterSelectOptions(q, table, col1, col2, col3, col4, col5, col6, col7, col8, col9)

	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	var v6s []T6
	var v7s []T7
	var v8s []T8
	var v9s []T9
	if err := q.Column(col1, col2, col3, col4, col5, col6, col7, col8, col9).
		Table(table).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s, &v6s, &v7s, &v8s, &v9s); err != nil {
		return nil, err
	}
	return handleTuple9Values(v1s, v2s, v3s, v4s, v5s, v6s, v7s, v8s, v9s), nil
}

func TypeTuple9SQL[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType, T8 ValueType, T9 ValueType](
	ctx context.Context, db RawIDB, query string, args ...any) ([]*Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], error) {
	var v1s []T1
	var v2s []T2
	var v3s []T3
	var v4s []T4
	var v5s []T5
	var v6s []T6
	var v7s []T7
	var v8s []T8
	var v9s []T9
	if err := db.NewRaw(query, args...).
		Scan(ctx, &v1s, &v2s, &v3s, &v4s, &v5s, &v6s, &v7s, &v8s, &v9s); err != nil {
		return nil, err
	}
	return handleTuple9Values(v1s, v2s, v3s, v4s, v5s, v6s, v7s, v8s, v9s), nil
}

func handleTuple9Values[T1 ValueType, T2 ValueType, T3 ValueType, T4 ValueType, T5 ValueType, T6 ValueType, T7 ValueType, T8 ValueType, T9 ValueType](
	v1s []T1, v2s []T2, v3s []T3, v4s []T4, v5s []T5, v6s []T6, v7s []T7, v8s []T8, v9s []T9) []*Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
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
	return ts
}

func setTableForSelectQuery(q *SelectQuery, table string) {
	if len(table) > 0 {
		q.Table(table)
	}
}

func setColumnForSelectQuery(q *SelectQuery, columns ...string) {
	for _, column := range columns {
		if len(column) > 0 {
			q.Column(column)
		}
	}
}

func afterSelectOptions(q *SelectQuery, table string, columns ...string) {
	setTableForSelectQuery(q, table)
	setColumnForSelectQuery(q, columns...)
}
