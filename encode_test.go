package flat_test

import (
	"github.com/rchowell/flat"
	"reflect"
	"testing"
	"time"
)

// simple struct
type s struct {
	A int
	B string
	C time.Time
	D []int
}

// nested struct
type ns struct {
	A int
	B string
	S s
}

// struct with pointers
type ps struct {
	A *int
	B *string
	C *time.Time
	D *[]int
	S *s
}

func TestFlatten(t *testing.T) {
	// use now for time comparison
	now := time.Now()

	// values to be used for pointer struct
	a := 0
	b := "hello"
	d := []int{4,5,6}

	tests := []struct {
		n string
		i interface{}
		e map[string]string
	}{
		{
			n: "integers",
			i: 1,
			e: map[string]string{"int": "1"},
		},
		{
			n: "strings",
			i: "hello",
			e: map[string]string{"string": "hello"},
		},
		{
			n: "slices/arrays",
			i: []string{"A", "B", "C"},
			e: map[string]string{
				"slice.0": "A",
				"slice.1": "B",
				"slice.2": "C",
			},
		},
		{
			n: "times",
			i: now,
			e: map[string]string{
				"time.Time": now.Format(time.RFC3339),
			},
		},
		{
			n: "structs",
			i: s{
				A: 0,
				B: "hello",
				C: now,
				D: []int{4,5,6},
			},
			e: map[string]string{
				"A": "0",
				"B": "hello",
				"C": now.Format(time.RFC3339),
				"D.0": "4",
				"D.1": "5",
				"D.2": "6",
			},
		},
		{
			n: "nested structs",
			i: ns{
				A: 0,
				B: "hello",
				S: s{
					A: 1,
					B: "goodbye",
					C: now,
					D: []int{4,5,6},
				},
			},
			e: map[string]string{
				"A": "0",
				"B": "hello",
				"S.A": "1",
				"S.B": "goodbye",
				"S.C": now.Format(time.RFC3339),
				"S.D.0": "4",
				"S.D.1": "5",
				"S.D.2": "6",
			},
		},
		{
			n: "structs with pointers",
			i: ps{
				A: &a,
				B: &b,
				C: &now,
				D: &d,
				S: &s{
					A: 1,
					B: "goodbye",
					C: now,
					D: []int{7,8,9},
				},
			},
			e: map[string]string{
				"A": "0",
				"B": "hello",
				"C": now.Format(time.RFC3339),
				"D.0": "4",
				"D.1": "5",
				"D.2": "6",
				"S.A": "1",
				"S.B": "goodbye",
				"S.C": now.Format(time.RFC3339),
				"S.D.0": "7",
				"S.D.1": "8",
				"S.D.2": "9",
			},
		},
		{
			n: "structs with nil values",
			i: s{
				A: 0,
				B: "",
				C: time.Time{},
				D: nil,
			},
			e: map[string]string{
				"A": "0",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.n, func(t *testing.T) {
			if got := flat.Flatten(tt.i); !reflect.DeepEqual(got, tt.e) {
				t.Errorf("Flatten() = %v, want %v", got, tt.e)
			}
		})
	}
}