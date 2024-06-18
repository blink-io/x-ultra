package x

import (
	"context"

	rdb "github.com/blink-io/x/sql/db"
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

func Insert[M Model](ctx context.Context, db rdb.RawIDB, m *M, ops ...InsertOption) error {
	q := db.NewInsert()
	opts := applyInsertOptions(ops...)
	handleInsertOptions(OperationInsert, q, opts)
	_, err := q.Model(m).Exec(ctx)
	return err
}

func BulkInsert[M Model](ctx context.Context, db rdb.RawIDB, ms []*M, ops ...InsertOption) error {
	q := db.NewInsert()
	o := applyInsertOptions(ops...)
	handleInsertOptions(OperationBulkInsert, q, o)
	_, err := q.Model(&ms).Exec(ctx)
	return err
}

func Update[M Model](ctx context.Context, db rdb.RawIDB, m *M, ops ...UpdateOption) error {
	q := db.NewUpdate()
	o := applyUpdateOptions(ops...)
	handleUpdateOptions(OperationUpdate, q, o)
	_, err := q.Model(m).
		WherePK().
		Exec(ctx)
	return err
}

func Delete[M Model, I ID](ctx context.Context, db rdb.RawIDB, ID I, field string, ops ...DeleteOption) error {
	q := db.NewDelete()
	o := applyDeleteOptions(ops...)
	handleDeleteOptions(OperationDelete, q, o)
	_, err := q.Model((*M)(nil)).
		Where("? = ?", rdb.Ident(field), ID).
		Exec(ctx)
	return err
}

func BulkDelete[M Model, I ID](ctx context.Context, db rdb.RawIDB, IDs []I, field string, ops ...DeleteOption) error {
	q := db.NewDelete()
	o := applyDeleteOptions(ops...)
	handleDeleteOptions(OperationBulkDelete, q, o)
	_, err := q.Model((*M)(nil)).
		Where("? IN (?)", rdb.Ident(field), rdb.In(IDs)).
		Exec(ctx)
	return err
}

func Get[M Model, I ID](ctx context.Context, db rdb.RawIDB, ID I, field string, ops ...SelectOption) (*M, error) {
	var m = new(M)
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectByPK, q, o)
	err := q.Model(m).
		Where("? = ?", rdb.Ident(field), ID).Limit(1).Scan(ctx, m)
	return m, err
}

func One[M Model](ctx context.Context, db rdb.RawIDB, ops ...SelectOption) (*M, error) {
	var m = new(M)
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectOne, q, o)
	err := q.Model(m).
		Limit(1).Scan(ctx, m)
	return m, err
}

func All[M Model](ctx context.Context, db rdb.RawIDB, ops ...SelectOption) ([]*M, error) {
	var ms []*M
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationSelectAll, q, o)
	err := q.Model(&ms).Scan(ctx)
	return ms, err
}

func Count[M Model](ctx context.Context, db rdb.RawIDB, ops ...SelectOption) (int, error) {
	q := db.NewSelect()
	o := applySelectOptions(ops...)
	handleSelectOptions(OperationCount, q, o)
	return q.Model((*M)(nil)).Count(ctx)
}

func Exists[M Model](ctx context.Context, db rdb.RawIDB, ops ...SelectOption) (bool, error) {
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
}

func handleDeleteOptions(op Operation, q *rdb.DeleteQuery, o *deleteOptions) {
	if q == nil || o == nil {
		return
	}
	if o.force {
		q.ForceDelete()
	}
	if r := o.returning; r != nil {
		q.Returning(r.Query, r.Args...)
	} else {
		q.Returning("NULL")
	}
}

func handleSelectOptions(op Operation, q *rdb.SelectQuery, o *selectOptions) {
	if q == nil || o == nil {
		return
	}
	if len(o.cols) > 0 {
		q.Column(o.cols...)
	}
	if len(o.where) > 0 {
		for _, w := range o.where {
			q.Where(w.Query, w.Args...)
		}
	}
}