package generics

import (
	"context"

	"github.com/uptrace/bun"
)

func Insert[M Model](ctx context.Context, db bun.IDB, m *M, ops ...InsertOption) error {
	q := db.NewInsert()
	for _, o := range ops {
		o(q)
	}
	_, err := q.Model(m).Exec(ctx)
	return err
}

func BulkInsert[M Model](ctx context.Context, db bun.IDB, ms []*M, ops ...InsertOption) error {
	q := db.NewInsert()
	for _, o := range ops {
		o(q)
	}
	_, err := q.Model(&ms).Exec(ctx)
	return err
}

func Update[M Model](ctx context.Context, db bun.IDB, m *M, ops ...UpdateOption) error {
	q := db.NewUpdate()
	for _, o := range ops {
		o(q)
	}
	_, err := q.Model(m).
		WherePK().
		Exec(ctx)
	return err
}

func Delete[M Model, I ID](ctx context.Context, db bun.IDB, ID I, field string, ops ...DeleteOption) error {
	q := db.NewDelete()
	for _, o := range ops {
		o(q)
	}
	_, err := q.
		Model((*M)(nil)).
		Where("?=?", bun.Ident(field), ID).
		Exec(ctx)
	return err
}

func BulkDelete[M Model, I ID](ctx context.Context, db bun.IDB, IDs []I, field string, ops ...DeleteOption) error {
	q := db.NewDelete()
	for _, o := range ops {
		o(q)
	}
	_, err := q.
		Model((*M)(nil)).
		Where("? IN (?)", bun.Ident(field), bun.In(IDs)).
		Exec(ctx)
	return err
}

func Get[M Model, I ID](ctx context.Context, db bun.IDB, ID I, field string, ops ...SelectOption) (*M, error) {
	var m = new(M)
	q := db.NewSelect()
	for _, o := range ops {
		o(q)
	}
	err := q.
		Model(m).
		Where("?=?", bun.Ident(field), ID).
		Limit(1).
		Scan(ctx, m)
	return m, err
}

func One[M Model](ctx context.Context, db bun.IDB, ops ...SelectOption) (*M, error) {
	var m = new(M)
	q := db.NewSelect()
	for _, o := range ops {
		o(q)
	}
	err := q.
		Model(m).
		Limit(1).
		Scan(ctx, m)
	return m, err
}

func All[M Model](ctx context.Context, db bun.IDB, ops ...SelectOption) ([]*M, error) {
	var ms []*M
	q := db.NewSelect()
	for _, o := range ops {
		o(q)
	}
	err := q.
		Model(&ms).
		Scan(ctx)
	return ms, err
}

func Count[M Model](ctx context.Context, db bun.IDB, ops ...SelectOption) (int, error) {
	q := db.NewSelect()
	for _, o := range ops {
		o(q)
	}
	return q.Model((*M)(nil)).Count(ctx)
}

func Exists[M Model](ctx context.Context, db bun.IDB, ops ...SelectOption) (bool, error) {
	q := db.NewSelect()
	for _, o := range ops {
		o(q)
	}
	return q.Model((*M)(nil)).Exists(ctx)
}
