package model

type IDModel struct {
	ID   int64  `bun:"id,pk,nullzero,autoincrement" db:"id,pk" json:"id,omitempty" toml:"id,omitempty" yaml:"id,omitempty" msgpack:"id,omitempty"`
	GUID string `bun:"guid,unique,notnull,type:varchar(60)" db:"guid" json:"guid,omitempty" toml:"guid,omitempty" yaml:"guid,omitempty" msgpack:"guid,omitempty"`
}
