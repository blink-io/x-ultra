package password_test

import (
	"fmt"

	"github.com/blink-io/x/password"
)

func ExampleCheckPassword() {
	valid, err := password.Check("admin", "pbkdf2_sha256$24000$JMO9TJawIXB1$5iz40fwwc+QW6lZY+TuNciua3YVMV3GXdgkhXrcvWag=")

	if valid {
		fmt.Println("Password is valid.")
	} else {
		if err == nil {
			fmt.Println("Password is valid.")
		} else {
			fmt.Printf("Error decoding password: %s\n", err)
		}
	}
}

func ExampleMakePassword() {
	hash, err := password.Make("my-password", password.GetRandomString(32), password.DefaultHasher)

	if err == nil {
		fmt.Println(hash)
	} else {
		fmt.Printf("Error encoding password: %s\n", err)
	}
}
