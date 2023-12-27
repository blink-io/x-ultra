package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/k0kubun/pp/v3"
	"github.com/stretchr/testify/require"
)

type User struct {
	name string
}

func TestSim_1(t *testing.T) {
	cn := func(ok bool) string {
		if ok {
			return "a"
		}
		return "b"
	}
	name := "a"
	name2 := cn(false)

	u1 := new(User)
	u1.name = name

	u2 := new(User)
	u2.name = name2

	require.Equal(t, u1, u2)
}

func TestIfaceStruct_1(t *testing.T) {
	type SS struct {
		Name string
	}
	var ss = SS(struct {
		Name string
	}{
		Name: "ok",
	})
	require.NotNil(t, ss)

	var ssp = &struct {
		Name string
	}{
		Name: "very ok",
	}
	var sspp = (*SS)(ssp)
	require.NotNil(t, sspp)
}

func autoincr() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestClosure(t *testing.T) {
	ac := autoincr()
	fmt.Println("i: ", ac())
	fmt.Println("i: ", ac())
	fmt.Println("i: ", ac())
	fmt.Println("i: ", ac())
	fmt.Println("i: ", ac())
}

type user struct {
	Name      string
	Birthdate time.Time
	Point     float64
	Valid     bool
}

func TestSprint(t *testing.T) {
	u := &user{
		Name:      "或绘画",
		Birthdate: time.Now(),
		Point:     3.1415,
		Valid:     true,
	}
	ss := fmt.Sprint(u)

	fmt.Println(ss)
}

func TestTime_Util(t *testing.T) {
	ntm := time.Now()
	fmt.Println("Time Now: ", ntm)

	var ztm time.Time
	require.Equal(t, false, ntm.IsZero())
	require.Equal(t, true, ztm.IsZero())
}

func TestColorPrint(t *testing.T) {
	m := map[string]string{"foo": "bar", "hello": "world"}
	pp.Print(m)
}
