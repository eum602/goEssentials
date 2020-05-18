package greet

import (
	"fmt"
	"testing"
)

//TestGreet tests if Greet returns a valid output
func TestGreet(t *testing.T) {
	s := Greet("Erick")
	if s != "Welcome Erick" {
		t.Error("got", s, "expected", "Welcome Erick")
	}
}

//ExampleGreet prints an example of use
func ExampleGreet() {
	fmt.Println(Greet("Erick"))
	//Output
	//Welcome Erick
}

//BenchmarkGreet tests performance by using "go test -bench Greet" or "go test -bench ."
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Erick")
	}
}
