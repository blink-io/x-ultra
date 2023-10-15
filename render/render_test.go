package render

import (
	"fmt"
	"strings"
	"testing"
	"time"

	_ "github.com/creasty/defaults"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type testJob struct {
	Title     string     `toml:"title" yaml:"title" msgpack:"title" cbor:"title"`
	Level     int        `toml:"level" yaml:"level" msgpack:"level" cbor:"level"`
	StartedAt *time.Time `toml:"started_at" yaml:"started_at" msgpack:"started_at" cbor:"started_at"`
}

type testUser struct {
	Name     string    `toml:"name" yaml:"name" msgpack:"name" cbor:"name"`
	Age      int       `toml:"age" yaml:"age" msgpack:"age" cbor:"age"`
	Location *string   `toml:"location" yaml:"location" msgpack:"location" cbor:"location"`
	JoinedAt time.Time `toml:"joined_at" yaml:"joined_at" msgpack:"joined_at" cbor:"joined_at"`
	Job      *testJob  `toml:"job" yaml:"job" msgpack:"job" cbor:"job"`
}

func createTestUser() *testUser {
	now := time.Now()
	loc := "广州"
	u1 := &testUser{
		Name:     "Hello你好",
		Age:      8,
		Location: &loc,
		JoinedAt: now,
		Job: &testJob{
			Title:     "Gooder",
			Level:     144,
			StartedAt: &now,
		},
	}
	return u1
}

func TestRender_Protobuf_1(t *testing.T) {
	r := New()
	var sb = new(strings.Builder)
	pb := timestamppb.Now()

	err := r.Protobuf(sb, 200, pb)
	require.NoError(t, err)

	fmt.Printf("sb:%#v\n", sb.String())
}

func TestRender_YAML_1(t *testing.T) {
	r := New()
	var sb = new(strings.Builder)

	err := r.YAML(sb, 200, createTestUser())
	require.NoError(t, err)

	fmt.Printf("sb:%s\n", sb.String())
}

func TestRender_TOML_1(t *testing.T) {
	r := New()
	var sb = new(strings.Builder)

	err := r.TOML(sb, 200, createTestUser())
	require.NoError(t, err)

	fmt.Printf("sb:%s\n", sb.String())
}

func TestRender_Msgpack_1(t *testing.T) {
	r := New()
	var sb = new(strings.Builder)

	err := r.Msgpack(sb, 200, createTestUser())
	require.NoError(t, err)

	fmt.Printf("sb:%s\n", sb.String())
}
