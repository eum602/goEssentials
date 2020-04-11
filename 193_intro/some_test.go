package main

import (
	"testing"
)

//TestSum tests sum function
func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	x := sum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x) //t.Error ==> equivalent to log followed by fail
	}
}

func TestSum1(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{ //table of tests
		test{[]int{1, 2}, 3},
		test{[]int{3, 2, 45}, 50},
	}

	for _, v := range tests {
		x := sum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}
