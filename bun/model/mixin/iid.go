package mixin

type IIDField struct {
	IID int64 `bun:"iid,type:integer,unique,notnull,autoincrement" db:"iid" json:"iid,omitempty" toml:"iid,omitempty" yaml:"iid,omitempty" msgpack:"iid,omitempty"`
}
