//go:build go1.10

package hooks

import "database/sql/driver"

func init() {
	interfaceTestCases = append(interfaceTestCases,
		struct {
			name               string
			expectedInterfaces []interface{}
		}{
			"ExecerQueryerContextSessionResetter", []interface{}{
				(*driver.ExecerContext)(nil),
				(*driver.QueryerContext)(nil),
				(*driver.SessionResetter)(nil)}})
}
