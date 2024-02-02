package id

import (
	"fmt"
	"testing"
	"time"

	"github.com/disgoorg/snowflake/v2"
	"github.com/stretchr/testify/require"
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

func TestSnowflake_1(t *testing.T) {
	id := snowflake.ID(123456789012345678)

	// deconstructs the snowflake ID into its component timestamp, worker ID, process ID, and increment
	id.Deconstruct()

	// the time.Time when the snowflake ID was generated
	id.Time()

	// the worker ID which the snowflake ID was generated
	id.WorkerID()

	// the process ID which the snowflake ID was generated
	id.ProcessID()

	// the sequence when the snowflake ID was generated
	id.Sequence()

	// returns the string representation of the snowflake ID
	id.String()

	// returns a new snowflake ID with worker ID, process ID, and sequence set to 0
	// this can be used for various pagination requests to the discord api
	tid := snowflake.New(time.Now())
	require.NotNil(t, tid)

	// returns a snowflake ID from an environment variable
	gid := snowflake.GetEnv("guild_id")
	require.NotNil(t, gid)

	// returns a snowflake ID from an environment variable and a bool indicating if the key was found
	ggid, found := snowflake.LookupEnv("guild_id")
	require.NotNil(t, ggid)
	require.NotNil(t, found)

	// returns the string as a snowflake ID or an error
	pid, err := snowflake.Parse("123456789012345678")
	require.NoError(t, err)
	require.NotNil(t, pid)

	// returns the string as a snowflake ID or panics if an error occurs
	mpid := snowflake.MustParse("123456789012345678")
	fmt.Println("MPID: ", mpid)
}
