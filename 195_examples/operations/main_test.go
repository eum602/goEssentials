package operations

import (
	"fmt"
	"testing"
)

//TestSum tests sum function. All examples that are shown for godoc must start with Example*
func ExampleSum() {
	fmt.Println(Sum(2, 3))
	//output:
	//5
}
//ExampleSum_1 is another example of Sum method; thus it must have ExampleSum_*. Eg. ExampleSum_first , ExampleSum_1
func ExampleSum_first() {
	fmt.Println(Sum(4, 3))
	//output:
	//7
}
func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	x := Sum(2, 3)
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
		x := Sum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}