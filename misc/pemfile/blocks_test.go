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
	"errors"
	"testing"

	"github.com/blink-io/x/misc/pemfile"
)

func TestReadBlock(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name     string
		filename string
		pemType  string
		err      error
	}{
		{
			name:     "RSAPrivateKey",
			pemType:  "RSA PRIVATE KEY",
			filename: "testdata/rsa_private_pkcs1.pem",
		},
		{
			name:     "RSAPublicKey/PKCS8",
			pemType:  "PUBLIC KEY",
			filename: "testdata/rsa_public_pkix.pem",
		},
		{
			name:     "ECPrivateKey",
			pemType:  "EC PRIVATE KEY",
			filename: "testdata/ec_private_sec1.pem",
		},
		{
			name:     "ECPublicKey",
			pemType:  "PUBLIC KEY",
			filename: "testdata/ec_public_pkix.pem",
		},
		{
			name:     "TrailingData",
			filename: "testdata/trailing_data.pem",
			err:      pemfile.ErrTrailingData,
		},
		{
			name:     "NotAPEMFile",
			filename: "testdata/not_a_pem.file",
			err:      pemfile.ErrNotFound,
		},
		{
			name:     "NoSuchFile",
			filename: "testdata/no_such_file.pem",
			err:      errors.New("no such file"),
		},
	}

	for _, tc := range testcases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			block, err := pemfile.ReadBlock(tc.filename)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if block == nil {
				if tc.pemType != "" {
					t.Errorf("expected PEM type %q", tc.pemType)
				}
			} else {
				if block.Type != tc.pemType {
					t.Errorf("got PEM type %q, want %q", block.Type, tc.pemType)
				}
			}
		})
	}
}

func TestReadBlocks(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name     string
		filename string
		length   int
		err      error
	}{
		{
			name:     "OneBlock",
			filename: "testdata/one_block.pem",
			length:   1,
		},
		{
			name:     "TwoBlocks",
			filename: "testdata/two_blocks.pem",
			length:   2,
		},
		{
			name:     "ThreeBlocks",
			filename: "testdata/three_blocks.pem",
			length:   3,
		},
		{
			name:     "EmptyFile",
			filename: "testdata/empty.file",
			length:   0,
			err:      errors.New("empty file"),
		},
		{
			name:     "TrailingData",
			filename: "testdata/trailing_data.pem",
			length:   0,
			err:      pemfile.ErrTrailingData,
		},
		{
			name:     "NotAPEMFile",
			filename: "testdata/not_a_pem.file",
			length:   0,
			err:      pemfile.ErrNotFound,
		},
		{
			name:     "NoSuchFile",
			filename: "testdata/no_such_file.pem",
			length:   0,
			err:      errors.New("no such file"),
		},
	}

	for _, tc := range testcases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			blocks, err := pemfile.ReadBlocks(tc.filename)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if len(blocks) != tc.length {
				t.Errorf("got %d blocks, want %d", len(blocks), tc.length)
			}
		})
	}
}
