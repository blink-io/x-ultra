package store

type EmptyStruct struct{}

var NilStruct = (*EmptyStruct)(nil)

type TokenMap map[string]*EmptyStruct

func (m TokenMap) Put(key string) {
	if m != nil {
		m[key] = NilStruct
	}
}

func NewTokenMap() TokenMap {
	return make(TokenMap)
}
