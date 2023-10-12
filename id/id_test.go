package id

import (
	"fmt"
	"testing"
)

func TestID_1(t *testing.T) {
	id1 := UUID()
	id2 := KSUID()
	id6 := ShortUUID()
	id7 := NanoID(20)
	id14 := Snowflake()
	id12 := XID()

	fmt.Println("UUID(), ", " id: ", id1, ", len: ", len(id1))
	fmt.Println("KSUID(), ", " id: ", id2, ", len: ", len(id2))
	fmt.Println("ShortUUID(), ", " id : ", id6, ", len: ", len(id6))
	fmt.Println("NanoID(), ", " id : ", id7, ", len: ", len(id7))
	fmt.Println("Snowflake(), ", " id: ", id14, ", len: ", len(id14))
	fmt.Println("XID(), ", " id: ", id12, ", len: ", len(id12))
}
