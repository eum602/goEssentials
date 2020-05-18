package mymath

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{34, 35}, 69},
		test{[]int{3, 3}, 6},
		test{[]int{1, 1, 1, 1, 4, 1}, 9},
	}

	for _, v := range tests {
		x := Sum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func ExampleSum() {
	fmt.Println(Sum(2, 3))
	//Output
	//5
}
