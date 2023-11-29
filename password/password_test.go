package password

import (
	"fmt"
	"strings"
	"testing"
)

func TestMakePasswordDefault(t *testing.T) {
	encoded, err := Make("admin", "1TMOT0Rohg3g", DefaultHasher)

	if err != nil {
		t.Fatalf("Make error: %s", err)
	}

	expected := "pbkdf2_sha256$216000$1TMOT0Rohg3g$N+wIigWW4zpxnFBwXTWK1Qt8C9aduBIAayDS2ee8KxI="

	if encoded != expected {
		t.Fatalf("Encoded hash %s does not match %s.", encoded, expected)
	}
}

func TestMakePasswordEmptySaltDefault(t *testing.T) {
	encoded, err := Make("admin", "", DefaultHasher)

	if err != nil {
		t.Fatalf("Make error: %s", err)
	}

	hasher := IdentifyHasher(encoded).String()

	if hasher != hasherToString(DefaultHasher) {
		t.Fatalf("hasher %s is not %s.", hasher, hasherToString(DefaultHasher))
	}
}

func TestMakePasswordArgon2Hasher(t *testing.T) {
	encoded, err := Make("admin", "", Argon2Hasher)

	if err != nil {
		t.Fatalf("Make password error: %s", err)
	}

	if !strings.HasPrefix(encoded, fmt.Sprintf("%s$", hasherToString(Argon2Hasher))) {
		t.Fatalf("Encoded password doesn't match algorithm (%s): %s", hasherToString(Argon2Hasher), encoded)
	}
}

func TestMakePasswordBCryptHasher(t *testing.T) {
	encoded, err := Make("admin", "", BCryptHasher)

	if err != nil {
		t.Fatalf("Make password error: %s", err)
	}

	if !strings.HasPrefix(encoded, fmt.Sprintf("%s$", hasherToString(BCryptHasher))) {
		t.Fatalf("Encoded password doesn't match algorithm (%s): %s", hasherToString(BCryptHasher), encoded)
	}
}

func TestMakePasswordBCryptSHA256Hasher(t *testing.T) {
	encoded, err := Make("admin", "", BCryptSHA256Hasher)

	if err != nil {
		t.Fatalf("Make password error: %s", err)
	}

	if !strings.HasPrefix(encoded, fmt.Sprintf("%s$", hasherToString(BCryptSHA256Hasher))) {
		t.Fatalf("Encoded password doesn't match algorithm (%s): %s", hasherToString(BCryptSHA256Hasher), encoded)
	}
}

func TestMakePasswordPBKDF2SHA256Hasher(t *testing.T) {
	encoded, err := Make("admin", "", PBKDF2SHA256Hasher)

	if err != nil {
		t.Fatalf("Make password error: %s", err)
	}

	if !strings.HasPrefix(encoded, fmt.Sprintf("%s$", hasherToString(PBKDF2SHA256Hasher))) {
		t.Fatalf("Encoded password doesn't match algorithm (%s): %s", hasherToString(PBKDF2SHA256Hasher), encoded)
	}
}

func TestCheckPasswordArgon2(t *testing.T) {
	valid, err := Check("admin", "argon2$argon2i$v=19$m=512,t=2,p=2$NnFZNGxmQTE1bmFV$kPPGrqD6dnRllcQeksFN+w")

	if err != nil {
		t.Fatalf("Check error: %s", err)
	}

	if !valid {
		t.Fatal("Password should be valid.")
	}
}

func TestCheckPasswordPBKDF2SHA256(t *testing.T) {
	valid, err := Check("admin", "pbkdf2_sha256$120000$WZrFZhpl3wOU$yPimyWN658IuAu0XErvg1Nowfd55k60hu4o+eDUlBDM=")

	if err != nil {
		t.Fatalf("Check error: %s", err)
	}

	if !valid {
		t.Fatal("Password should be valid.")
	}
}

func TestCheckPasswordBCrypto(t *testing.T) {
	valid, err := Check("admin", "bcrypt$$2b$12$qcNExitVe89wMG.nmRD4Qupn2hFm0pxvnu6VC.w6LShOx30l.F9/.")

	if err != nil {
		t.Fatalf("Check error: %s", err)
	}

	if !valid {
		t.Fatal("Password should be valid.")
	}
}

func TestCheckPasswordBCryptoSHA256(t *testing.T) {
	valid, err := Check("admin", "bcrypt_sha256$$2b$12$WZK9cb9qKN.Q5LCYPq/gj.6gvry1b37HUsJER6KhQBnIWmPyyaaqi")

	if err != nil {
		t.Fatalf("Check error: %s", err)
	}

	if !valid {
		t.Fatal("Password should be valid.")
	}
}

func TestIsPasswordUsableWithValidPassword(t *testing.T) {
	usable := IsUsable("pbkdf2_sha256$24000$JMO9TJawIXB1$5iz40fwwc+QW6lZY+TuNciua3YVMV3GXdgkhXrcvWag=")

	if !usable {
		t.Fatal("Password should be usable.")
	}
}

func TestIsPasswordUsableWithUnusablePassword1(t *testing.T) {
	usable := IsUsable("!")

	if usable {
		t.Fatal("Password should be unusable.")
	}
}

func TestIsPasswordUsableWithUnusablePassword2(t *testing.T) {
	usable := IsUsable("!password")

	if usable {
		t.Fatal("Password should be unusable.")
	}
}

func TestIsPasswordUsableWithEmptyPassword(t *testing.T) {
	usable := IsUsable("")

	if usable {
		t.Fatal("Password should be unusable.")
	}
}
