package x

import (
	"context"

	rdb "github.com/blink-io/x/bun"
)

type Operation int

const (
	OperationInsert = iota
	OperationBulkInsert
	OperationUpdate
	OperationDelete
	OperationBulkDelete
	OperationSelectOne
	OperationSelectByPK
	OperationSelectAll
	OperationCount
	OperationExists
)

type IDSlice[I IDType] []I

type ModelSlice[M ModelType] []*M

func (s ModelSlice[I]) Emtpy() bool {
	return len(s) == 0
}

// Insert executes insert by a given model.
func Insert[M ModelType](ctx context.Context, db rdb.RawIDB, m *M, ops ...InsertOption) error {
	q := db.NewInsert()
	opts := applyInsertOptions(ops...)
	handleInsertOptions(OperationInsert, q, opts)
	_, err := q.Model(m).Exec(ctx)
	return err
}

// BulkInsert executes bulk insert by given model slice.
func BulkInsert[M ModelType](ctx context.Context, db rdb.RawIDB, ms ModelSlice[M], ops ...InsertOption) error {
	q := db.NewInsert()
	o := applyInsertOptions(ops...)
	handleInsertOptions(OperationBulkInsert, q, o)
	_, err := q.Model(&ms).Exec(ctx)
	return err
}

// Update updates a record by a given model with PK.
func Update[M ModelType](ctx context.Context, db rdb.RawIDB, m *M, ops ...UpdateOption) error {
	q := db.NewUpdate()
	o := applyUpdateOptions(ops...)
	handleUpdateOptions(OperationUpdate, q, o)
	_, err := q.Model(m).
		WherePK().
		Exec(ctx)
	return err
}

// Delete deletes a record by a given model with a specified column.
func Delete[M ModelType, I IDType](ctx context.Context, db rdb.RawIDB, ID I, column string, ops ...DeleteOption) error {
	q := db.NewDelete()
	o := applyDeleteOptions(ops...)
	handleDeleteOptions(OperationDelete, q, o)
	_, err := q.Model((*M)(nil)).
		Where("? = ?", rdb.Ident(column), ID).
		Exec(ctx)
	return err
}

// BulkDelete deletes records by a given model with a specified column.
func BulkDelete[M ModelType, I IDType](ctx context.Context, db rdb.RawIDB, IDs IDSlice[I], column string, ops ...DeleteOption) error {
	q := db.NewDelete()
	o := applyDeleteOptions(ops...)
	handleDeleteOptions(OperationBulkDelete, q, o)
	_, err := q.Model((*M)(nil)).
		Where("? IN (?)", rdb.Ident(column), rdb.In(IDs)).
		Exec(ctx)
	return err
}

// Get gets a record by a given model with a specified column and ID value.
func Get[M ModelType, I IDType](ctx context.Context, db rdb.RawIDB, ID I, column string, ops ...SelectOption) (*M, error) {
	var m = new(M)
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectByPK, q, o)
	err := q.Model(m).
		Where("? = ?", rdb.Ident(column), ID).Limit(1).
		Scan(ctx, m)
	return m, err
}

func One[M ModelType](ctx context.Context, db rdb.RawIDB, ops ...SelectOption) (*M, error) {
	m := new(M)
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectOne, q, o)
	err := q.Model(m).
		Limit(1).
		Scan(ctx, m)
	return m, err
}

func All[M ModelType](ctx context.Context, db rdb.RawIDB, ops ...SelectOption) (ModelSlice[M], error) {
	var ms ModelSlice[M]
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)
	err := q.Model(&ms).
		Scan(ctx)
	return ms, err
}

func Count[M ModelType](ctx context.Context, db rdb.RawIDB, ops ...SelectOption) (int, error) {
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationCount, q, o)
	return q.Model((*M)(nil)).Count(ctx)
}

func Exists[M ModelType](ctx context.Context, db rdb.RawIDB, ops ...SelectOption) (bool, error) {
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationExists, q, o)
	return q.Model((*M)(nil)).Exists(ctx)
}

func handleInsertOptions(op Operation, q *rdb.InsertQuery, o *insertOptions) {
	if q == nil || o == nil {
		return
	}
	if r := o.returning; r != nil {
		q.Returning(r.Query, r.Args...)
	} else {
		q.Returning("NULL")
	}
	if o.ignore {
		q.Ignore()
	}
	if len(o.columns) > 0 {
		q.Column(o.columns...)
	}
	if len(o.columnExprs) > 0 {
		for _, c := range o.columnExprs {
			q.ColumnExpr(c.Query, c.Args...)
		}
	}
	if len(o.excludeColumns) > 0 {
		q.ExcludeColumn(o.excludeColumns...)
	}
}

func handleUpdateOptions(op Operation, q *rdb.UpdateQuery, o *updateOptions) {
	if q == nil || o == nil {
		return
	}
	if o.omitZero {
		q.OmitZero()
	}
	if r := o.returning; r != nil {
		q.Returning(r.Query, r.Args...)
	} else {
		q.Returning("NULL")
	}
	if o.bulk {
		q.Bulk()
	}
}

func handleDeleteOptions(op Operation, q *rdb.DeleteQuery, o *deleteOptions) {
	if q == nil || o == nil {
		return
	}
	if o.forceDelete {
		q.ForceDelete()
	}
	if r := o.returning; r != nil {
		q.Returning(r.Query, r.Args...)
	} else {
		q.Returning("NULL")
	}
	if o.where != nil {
		for _, w := range o.where {
			q.Where(w.Query, w.Args...)
		}
	}
	if o.whereOr != nil {
		for _, w := range o.where {
			q.WhereOr(w.Query, w.Args...)
		}
	}
}

func handleSelectOptions(op Operation, q *rdb.SelectQuery, o *selectOptions) {
	if q == nil || o == nil {
		return
	}

	if len(o.columns) > 0 {
		q.Column(o.columns...)
	}
	if len(o.where) > 0 {
		for _, w := range o.where {
			q.Where(w.Query, w.Args...)
		}
	}
	if len(o.columns) > 0 {
		q.Column(o.columns...)
	}
	if len(o.columnExprs) > 0 {
		for _, c := range o.columnExprs {
			q.ColumnExpr(c.Query, c.Args...)
		}
	}
	if len(o.excludeColumns) > 0 {
		q.ExcludeColumn(o.excludeColumns...)
	}
	if op == OperationSelectAll {
		if o.queryFunc != nil {
			q.Apply(o.queryFunc)
		}
		if o.distinct {
			q.Distinct()
		}
		if len(o.distinctOn) > 0 {
			for _, d := range o.distinctOn {
				q.DistinctOn(d.Query, d.Args...)
			}
		}
		if o.limit > 0 {
			q.Limit(o.limit)
		}
		if o.offset > 0 {
			q.Offset(o.offset)
		}
	}
}
