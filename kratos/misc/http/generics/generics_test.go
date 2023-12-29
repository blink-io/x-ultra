package generics

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGenerics_Handle_1(t *testing.T) {
	methods := []string{
		http.MethodGet,
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions,
		http.MethodConnect,
		http.MethodTrace,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
	}
	fn := func(m string) {
		switch m {
		case http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch:
			fmt.Println("Method: ", m)
		default:
		}
	}
	for _, m := range methods {
		fn(m)
	}
}
