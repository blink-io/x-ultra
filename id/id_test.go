package id

import (
	"fmt"
	"testing"
)

func TestID_1(t *testing.T) {
	uuidv := UUID()
	uuidv4v := UUIDV4()
	ksuidv := KSUID()
	shortiuuidv := ShortUUID()
	nanoidv := NanoID(20)
	sfv := Snowflake()
	xidv := XID()
	shortidv := ShortID()
	ulidv := ULID()
	guidv := GUID()

	fmt.Println("UUID(),        len: ", len(uuidv), " id: ", uuidv)
	fmt.Println("UUIDV4(),      len: ", len(uuidv4v), " id: ", uuidv4v)
	fmt.Println("KSUID(),       len: ", len(ksuidv), " id: ", ksuidv)
	fmt.Println("NanoID(),      len: ", len(nanoidv), " id: ", nanoidv)
	fmt.Println("Snowflake(),   len: ", len(sfv), " id: ", sfv)
	fmt.Println("ShortID(),     len: ", len(shortidv), " id: ", shortidv)
	fmt.Println("ShortUUID(),   len: ", len(shortiuuidv), " id: ", shortiuuidv)
	fmt.Println("XID(),         len: ", len(xidv), " id: ", xidv)
	fmt.Println("ULID(),        len: ", len(ulidv), " id: ", ulidv)
	fmt.Println("GUID(),        len: ", len(guidv), " id: ", guidv)
}
