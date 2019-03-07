package adder_test

import (
	"testing"

	"github.com/ptiger10/go-module-example/adder"
)

func TestAdder(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 4}, 5},
		{[]int{5}, 5},
	}
	for _, test := range tests {
		if got := adder.Add(test.input...); got != test.want {
			t.Errorf("Add(%v) returned %d, want %d",
				test.input, got, test.want)
		}
	}
}
