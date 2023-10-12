package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ALL(t *testing.T) {
	TestUsername_valid_1(t)
	TestUsername_invalid_1(t)
}

func TestUsername_valid_1(t *testing.T) {
	validUsernames :=
		[]string{
			"user_name",
			"u18n_abc",
			"username_",
			"userName_",
			"Uname",
			"U_name_bcd_12",
			"U_name____",
		}

	for _, n := range validUsernames {
		valid := usernameRegex.MatchString(n)
		require.Equalf(t, true, valid, "invalid for : %s", n)
	}
}

func TestUsername_invalid_1(t *testing.T) {
	invalidUsernames :=
		[]string{
			"_name188",
			"18good",
			"___abc",
			"_1_n_2_",
			"_Uname",
			"18_an_18",
		}

	for _, n := range invalidUsernames {
		valid := usernameRegex.MatchString(n)
		require.Equalf(t, false, valid, "valid for : %s", n)
	}
}
