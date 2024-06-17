/*
Copyright (c) 2020 GMO GlobalSign, Inc.

Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at

https://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pemfile_test

import (
	"bytes"
	"encoding/pem"
	"errors"
	"reflect"
	"testing"

	"github.com/blink-io/x/misc/pemfile"
)

func TestReadCSR(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		filename string
		err      error
	}{
		{
			filename: "testdata/example_csr.pem",
		},
		{
			filename: "testdata/no_such.file",
			err:      errors.New("no such file"),
		},
		{
			filename: "testdata/example_root_ca.pem",
			err:      errors.New("not a CSR"),
		},
	}

	for _, tc := range testcases {
		tc := tc

		t.Run(tc.filename, func(t *testing.T) {
			t.Parallel()

			_, err := pemfile.ReadCSR(tc.filename)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}
		})
	}
}

func TestWriteCSR(t *testing.T) {
	t.Parallel()

	var testcases = []string{
		"testdata/example_csr.pem",
	}

	for _, tc := range testcases {
		tc := tc

		t.Run(tc, func(t *testing.T) {
			t.Parallel()

			block, err := pemfile.ReadBlock(tc)
			if err != nil {
				t.Fatalf("failed to read PEM block: %v", err)
			}

			csr, err := pemfile.ReadCSR(tc)
			if err != nil {
				t.Fatalf("failed to read certificate request: %v", err)
			}

			buffer := bytes.NewBuffer([]byte{})
			err = pemfile.WriteCSR(buffer, csr)
			if err != nil {
				t.Fatalf("failed to write certificate request: %v", err)
			}

			got, _ := pem.Decode(buffer.Bytes())
			if !reflect.DeepEqual(got, block) {
				t.Fatalf("got %v, want %v", got, block)
			}
		})
	}
}
