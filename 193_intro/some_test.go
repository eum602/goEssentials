package main

import (
	"testing"
)

//TestSum tests sum function
func TestSum(t *testing.T) {
	x := sum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x) //t.Error ==> equivalent to log followed by fail
	}
}
