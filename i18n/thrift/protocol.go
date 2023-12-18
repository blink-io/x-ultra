package thrift

type Protocol int

const (
	Binary Protocol = iota
	Compact
	SimpleJSON
	JSON
)
