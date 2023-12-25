package model

type Meta[T any, C any] struct {
	PK        string
	Label     string
	TableName string
	Alias     string
	Type      *T
	Columns   *C
}

const (
	FieldNameCreatedAt = "created_at"
	FieldNameUpdatedAt = "updated_at"
	FieldNameDeletedAt = "deleted_at"

	FieldNameCreateBy  = "created_by"
	FieldNameUpdatedBy = "updated_by"
	FieldNameDeletedBy = "deleted_by"

	FieldNameIsDeleted = "is_deleted"
)
