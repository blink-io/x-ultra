package model

type IDModel struct {
	ID   int64  `bun:"id,pk,autoincrement" db:"id,pk" goqu:"skipinsert" json:"id,omitempty" toml:"id,omitempty" yaml:"id,omitempty" msgpack:"id,omitempty"`
	GUID string `bun:"guid,type:varchar(60),unique,notnull" db:"guid" json:"guid,omitempty" toml:"guid,omitempty" yaml:"guid,omitempty" msgpack:"guid,omitempty"`
}