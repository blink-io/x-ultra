package mixin

type IDField struct {
	ID string `bun:"id,pk,type:varchar(60)" db:"id" json:"id,omitempty" toml:"id,omitempty" yaml:"id,omitempty" msgpack:"id,omitempty"`
}
