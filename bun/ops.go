package bun

import (
	"context"
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

func (it IDSlice[I]) Emtpy() bool {
	return len(it) == 0
}

type ModelSlice[M ModelType] []*M

func (st ModelSlice[I]) Emtpy() bool {
	return len(st) == 0
}

// DoInsert executes insert by a given model.
func DoInsert[M ModelType](ctx context.Context, db RawIDB, m *M, ops ...DoInsertOption) error {
	q := db.NewInsert()
	opts := applyDoInsertOptions(ops...)
	handleDoInsertOptions(OperationInsert, q, opts)
	_, err := q.Model(m).Exec(ctx)
	return err
}

// DoBulkInsert executes bulk insert by given model slice.
func DoBulkInsert[M ModelType](ctx context.Context, db RawIDB, ms ModelSlice[M], ops ...DoInsertOption) error {
	q := db.NewInsert()
	o := applyDoInsertOptions(ops...)
	handleDoInsertOptions(OperationBulkInsert, q, o)
	_, err := q.Model(&ms).Exec(ctx)
	return err
}

// DoUpdate updates a record by a given model with PK.
func DoUpdate[M ModelType](ctx context.Context, db RawIDB, m *M, ops ...DoUpdateOption) error {
	q := db.NewUpdate()
	o := applyDoUpdateOptions(ops...)
	handleDoUpdateOptions(OperationUpdate, q, o)
	_, err := q.Model(m).
		WherePK().
		Exec(ctx)
	return err
}

// DoDelete deletes a record by a given model with a specified column.
func DoDelete[M ModelType, I IDType](ctx context.Context, db RawIDB, ID I, column string, ops ...DoDeleteOption) error {
	q := db.NewDelete()
	o := applyDoDeleteOptions(ops...)
	handleDoDeleteOptions(OperationDelete, q, o)
	_, err := q.Model((*M)(nil)).
		Where("? = ?", Ident(column), ID).
		Exec(ctx)
	return err
}

// DoBulkDelete deletes records by a given model with a specified column.
func DoBulkDelete[M ModelType, I IDType](ctx context.Context, db RawIDB, IDs IDSlice[I], column string, ops ...DoDeleteOption) error {
	q := db.NewDelete()
	o := applyDoDeleteOptions(ops...)
	handleDoDeleteOptions(OperationBulkDelete, q, o)
	_, err := q.Model((*M)(nil)).
		Where("? IN (?)", Ident(column), In(IDs)).
		Exec(ctx)
	return err
}

// DoGet gets a record by a given model with a specified column and I value.
func DoGet[M ModelType, I IDType](ctx context.Context, db RawIDB, ID I, column string, ops ...DoSelectOption) (*M, error) {
	var m = new(M)
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectByPK, q, o)
	err := q.Model(m).
		Where("? = ?", Ident(column), ID).Limit(1).
		Scan(ctx, m)
	return m, err
}

func DoOne[M ModelType](ctx context.Context, db RawIDB, ops ...DoSelectOption) (*M, error) {
	m := new(M)
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectOne, q, o)
	err := q.Model(m).
		Limit(1).
		Scan(ctx, m)
	return m, err
}

func DoAll[M ModelType](ctx context.Context, db RawIDB, ops ...DoSelectOption) (ModelSlice[M], error) {
	var ms ModelSlice[M]
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationSelectAll, q, o)
	err := q.Model(&ms).
		Scan(ctx)
	return ms, err
}

func DoCount[M ModelType](ctx context.Context, db RawIDB, ops ...DoSelectOption) (int, error) {
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationCount, q, o)
	return q.Model((*M)(nil)).Count(ctx)
}

func DoExists[M ModelType](ctx context.Context, db RawIDB, ops ...DoSelectOption) (bool, error) {
	q := db.NewSelect()
	o := applyDoSelectOptions(ops...)
	handleDoSelectOptions(OperationExists, q, o)
	return q.Model((*M)(nil)).Exists(ctx)
}

func handleDoInsertOptions(op Operation, q *InsertQuery, o *doInsertOptions) {
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

func handleDoUpdateOptions(op Operation, q *UpdateQuery, o *doUpdateOptions) {
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

func handleDoDeleteOptions(op Operation, q *DeleteQuery, o *doDeleteOptions) {
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

func handleDoSelectOptions(op Operation, q *SelectQuery, o *doSelectOptions) {
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
		if o.queryBuilder != nil {
			q.ApplyQueryBuilder(o.queryBuilder)
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
	} else if op == OperationExists || op == OperationCount {
		if o.queryFunc != nil {
			q.Apply(o.queryFunc)
		}
		if o.queryBuilder != nil {
			q.ApplyQueryBuilder(o.queryBuilder)
		}
	}
}
