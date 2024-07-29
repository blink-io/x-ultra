package bun

type (
	doInsertOptions struct {
		ignore         bool
		columns        []string
		excludeColumns []string
		columnExprs    []*QueryWithArgs
		returning      *QueryWithArgs
	}

	DoInsertOption func(*doInsertOptions)

	doUpdateOptions struct {
		omitZero       bool
		bulk           bool
		columns        []string
		excludeColumns []string
		FQN            string
		forceIndexes   []string
		ignoreIndexes  []string
		modelTableExpr *QueryWithArgs
		returning      *QueryWithArgs
	}

	DoUpdateOption func(*doUpdateOptions)

	doSelectOptions struct {
		queryFunc      func(*SelectQuery) *SelectQuery
		queryBuilder   func(QueryBuilder) QueryBuilder
		distinct       bool
		distinctOn     []*QueryWithArgs
		limit          int
		offset         int
		columns        []string
		excludeColumns []string
		columnExprs    []*QueryWithArgs
		orders         []string
		where          []*QueryWithArgs
		whereOr        []*QueryWithArgs
	}

	DoSelectOption func(*doSelectOptions)

	doDeleteOptions struct {
		forceDelete bool
		returning   *QueryWithArgs
		where       []*QueryWithArgs
		whereOr     []*QueryWithArgs
	}

	DoDeleteOption func(*doDeleteOptions)
)

func applyDoInsertOptions(ops ...DoInsertOption) *doInsertOptions {
	opts := new(doInsertOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func DoWithInsertIgnore() DoInsertOption {
	return func(o *doInsertOptions) {
		o.ignore = true
	}
}

func DoWithInsertReturning(query string, args ...any) DoInsertOption {
	return func(o *doInsertOptions) {
		o.returning = doSafeQuery(query, args...)
	}
}

func applyDoUpdateOptions(ops ...DoUpdateOption) *doUpdateOptions {
	opts := new(doUpdateOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func DoWithUpdateOmitZero() DoUpdateOption {
	return func(o *doUpdateOptions) {
		o.omitZero = true
	}
}

func DoWithUpdateReturning(query string, args ...any) DoUpdateOption {
	return func(o *doUpdateOptions) {
		o.returning = doSafeQuery(query, args...)
	}
}

func applyDoDeleteOptions(ops ...DoDeleteOption) *doDeleteOptions {
	opts := new(doDeleteOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func DoWithDeleteForce() DoDeleteOption {
	return func(o *doDeleteOptions) {
		o.forceDelete = true
	}
}

func DoWithDeleteReturning(query string, args ...any) DoDeleteOption {
	return func(o *doDeleteOptions) {
		o.returning = doSafeQuery(query, args...)
	}
}

func DoWithDeleteWhere(query string, args ...any) DoDeleteOption {
	return func(o *doDeleteOptions) {
		o.where = append(
			o.where,
			doSafeQuery(query, args...),
		)
	}
}

func DoWithDeleteWhereOr(query string, args ...any) DoDeleteOption {
	return func(o *doDeleteOptions) {
		o.whereOr = append(
			o.whereOr,
			doSafeQuery(query, args...),
		)
	}
}

func applyDoSelectOptions(ops ...DoSelectOption) *doSelectOptions {
	opts := new(doSelectOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func DoWithSelectLimit(limit int) DoSelectOption {
	return func(o *doSelectOptions) {
		o.limit = limit
	}
}

func DoWithSelectOffset(offset int) DoSelectOption {
	return func(o *doSelectOptions) {
		o.offset = offset
	}
}

func DoWithSelectWhere(query string, args ...any) DoSelectOption {
	return func(o *doSelectOptions) {
		o.where = append(
			o.where,
			doSafeQuery(query, args...),
		)
	}
}

func DoWithSelectWhereOr(query string, args ...any) DoSelectOption {
	return func(o *doSelectOptions) {
		o.whereOr = append(
			o.whereOr,
			doSafeQuery(query, args...),
		)
	}
}

func DoWithSelectColumns(columns ...string) DoSelectOption {
	return func(o *doSelectOptions) {
		o.columns = append(o.columns, columns...)
	}
}

func DoWithSelectOrders(orders ...string) DoSelectOption {
	return func(o *doSelectOptions) {
		o.orders = append(o.orders, orders...)
	}
}

func DoWithSelectQuery(queryFunc func(*SelectQuery) *SelectQuery) DoSelectOption {
	return func(o *doSelectOptions) {
		o.queryFunc = queryFunc
	}
}

func doSafeQuery(query string, args ...any) *QueryWithArgs {
	sq := SafeQuery(query, args)
	return &sq
}
