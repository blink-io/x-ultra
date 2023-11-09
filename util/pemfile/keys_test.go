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
	"crypto/ecdsa"
	"crypto/rsa"
	"errors"
	"reflect"
	"testing"

	"github.com/blink-io/x/util/pemfile"
)

func TestReadPrivateKey(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		filename string
		keyType  reflect.Type
		err      error
	}{
		{
			filename: "testdata/rsa_private_pkcs1.pem",
			keyType:  reflect.TypeOf((*rsa.PrivateKey)(nil)),
		},
		{
			filename: "testdata/rsa_private_pkcs8.pem",
			keyType:  reflect.TypeOf((*rsa.PrivateKey)(nil)),
		},
		{
			filename: "testdata/ec_private_sec1.pem",
			keyType:  reflect.TypeOf((*ecdsa.PrivateKey)(nil)),
		},
		{
			filename: "testdata/ec_private_pkcs8.pem",
			keyType:  reflect.TypeOf((*ecdsa.PrivateKey)(nil)),
		},
		{
			filename: "testdata/rsa_public_pkix.pem",
			err:      errors.New("not a private key"),
		},
		{
			filename: "testdata/no_such_file.pem",
			err:      errors.New("no such file"),
		},
	}

	for _, tc := range testcases {
		tc := tc

		t.Run(tc.filename, func(t *testing.T) {
			t.Parallel()

			key, err := pemfile.ReadPrivateKey(tc.filename)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if keyType := reflect.TypeOf(key); !reflect.DeepEqual(keyType, tc.keyType) {
				t.Errorf("got key type %v, want %v", keyType, tc.keyType)
			}
		})
	}
}

func TestReadPrivateKeyWithPasswordFunc(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		filename string
		keyType  reflect.Type
		pwfunc   func(string, string) ([]byte, error)
		err      error
	}{
		{
			filename: "testdata/rsa_private_pkcs1_encrypted.pem",
			keyType:  reflect.TypeOf((*rsa.PrivateKey)(nil)),
			pwfunc:   func(string, string) ([]byte, error) { return []byte("strongpassword"), nil },
		},
		{
			filename: "testdata/rsa_private_pkcs1_encrypted.pem",
			pwfunc:   func(string, string) ([]byte, error) { return []byte("wrongpassword"), nil },
			err:      errors.New("decryption failure"),
		},
		{
			filename: "testdata/rsa_private_pkcs1_encrypted.pem",
			pwfunc:   func(string, string) ([]byte, error) { return nil, errors.New("error") },
			err:      errors.New("password failure"),
		},
		{
			filename: "testdata/rsa_private_pkcs1.pem",
			keyType:  reflect.TypeOf((*rsa.PrivateKey)(nil)),
		},
		{
			filename: "testdata/rsa_private_pkcs8.pem",
			keyType:  reflect.TypeOf((*rsa.PrivateKey)(nil)),
		},
		{
			filename: "testdata/ec_private_sec1.pem",
			keyType:  reflect.TypeOf((*ecdsa.PrivateKey)(nil)),
		},
		{
			filename: "testdata/ec_private_pkcs8.pem",
			keyType:  reflect.TypeOf((*ecdsa.PrivateKey)(nil)),
		},
		{
			filename: "testdata/rsa_public_pkix.pem",
			err:      errors.New("not a private key"),
		},
		{
			filename: "testdata/no_such_file.pem",
			err:      errors.New("no such file"),
		},
	}

	for _, tc := range testcases {
		tc := tc

		t.Run(tc.filename, func(t *testing.T) {
			t.Parallel()

			key, err := pemfile.ReadPrivateKeyWithPasswordFunc(tc.filename, tc.pwfunc)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if keyType := reflect.TypeOf(key); !reflect.DeepEqual(keyType, tc.keyType) {
				t.Errorf("got key type %v, want %v", keyType, tc.keyType)
			}
		})
	}
}

func TestReadPublicKey(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		filename string
		keyType  reflect.Type
		err      error
	}{
		{
			filename: "testdata/rsa_public_pkix.pem",
			keyType:  reflect.TypeOf((*rsa.PublicKey)(nil)),
		},
		{
			filename: "testdata/rsa_public_pkcs1.pem",
			keyType:  reflect.TypeOf((*rsa.PublicKey)(nil)),
		},
		{
			filename: "testdata/ec_public_pkix.pem",
			keyType:  reflect.TypeOf((*ecdsa.PublicKey)(nil)),
		},
		{
			filename: "testdata/rsa_private_pkcs1.pem",
			err:      errors.New("not a public key"),
		},
		{
			filename: "testdata/no_such_file.pem",
			err:      errors.New("no such file"),
		},
	}

	for _, tc := range testcases {
		tc := tc

		t.Run(tc.filename, func(t *testing.T) {
			t.Parallel()

			key, err := pemfile.ReadPublicKey(tc.filename)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if keyType := reflect.TypeOf(key); !reflect.DeepEqual(keyType, tc.keyType) {
				t.Errorf("got key type %v, want %v", keyType, tc.keyType)
			}
		})
	}
}
