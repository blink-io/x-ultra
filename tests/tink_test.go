package tests

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tink-crypto/tink-go/v2/aead"
	"github.com/tink-crypto/tink-go/v2/insecurecleartextkeyset"
	"github.com/tink-crypto/tink-go/v2/keyset"
)

func TestTink_1(t *testing.T) {
	// A keyset created with "tinkey create-keyset --key-template=AES256_GCM". Note
	// that this keyset has the secret key information in cleartext.
	jsonKeyset := `{
                        "key": [{
                                        "keyData": {
                                                        "keyMaterialType":
                                                                        "SYMMETRIC",
                                                        "typeUrl":
                                                                        "type.googleapis.com/google.crypto.tink.AesGcmKey",
                                                        "value":
                                                                        "GiBWyUfGgYk3RTRhj/LIUzSudIWlyjCftCOypTr0jCNSLg=="
                                        },
                                        "keyId": 294406504,
                                        "outputPrefixType": "TINK",
                                        "status": "ENABLED"
                        }],
                        "primaryKeyId": 294406504
        }`

	// Create a keyset handle from the cleartext keyset in the previous
	// step. The keyset handle provides abstract access to the underlying keyset to
	// limit the exposure of accessing the raw key material. WARNING: In practice,
	// it is unlikely you will want to use a insecurecleartextkeyset, as it implies
	// that your key material is passed in cleartext, which is a security risk.
	// Consider encrypting it with a remote key in Cloud KMS, AWS KMS or HashiCorp Vault.
	// See https://github.com/google/tink/blob/master/docs/GOLANG-HOWTO.md#storing-and-loading-existing-keysets.
	keysetHandle, err := insecurecleartextkeyset.Read(
		keyset.NewJSONReader(bytes.NewBufferString(jsonKeyset)))
	require.NoError(t, err)

	// Retrieve the AEAD primitive we want to use from the keyset handle.
	primitive, err := aead.New(keysetHandle)
	require.NoError(t, err)

	// Use the primitive to encrypt a message. In this case the primary key of the
	// keyset will be used (which is also the only key in this example).
	plaintext := []byte("message")
	associatedData := []byte("associated data")
	ciphertext, err := primitive.Encrypt(plaintext, associatedData)
	require.NoError(t, err)

	fmt.Println("ciphertext in base64:  ", base64.StdEncoding.EncodeToString(ciphertext))
	// Use the primitive to decrypt the message. Decrypt finds the correct key in
	// the keyset and decrypts the ciphertext. If no key is found or decryption
	// fails, it returns an error.
	decrypted, err := primitive.Decrypt(ciphertext, associatedData)
	require.NoError(t, err)

	fmt.Println(string(decrypted))
	// Output: message
}
