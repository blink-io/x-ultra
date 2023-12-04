package cache

type Codec interface {
	Encode(v any) ([]byte, error)
	Decode([]byte) (v any, err error)
}
