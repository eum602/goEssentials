package main

import (
	"fmt"
)

func main() {
	definitions()
	example1()
}

func definitions() {
	fmt.Println(`
	Maps are key/value store, store some values based upon some keys. Then it allows super fast and super
	efficient lookup.
	*Maps are unordered list.

	* m := map[string]int{ //key is string and value is int. map[string]int ==> it is the type.
		"James":  32,
		"eum602": 31,
	}`)
}

func example1() {
	fmt.Println("Example 1")
	m := map[string]int{ //key is string and value is int. map[string]int ==> it is the type.
		"James":  32,
		"eum602": 31,
	}
	fmt.Println(m)
	fmt.Println(`m["eum602"] ==>`, m["eum602"])                                           //31
	fmt.Println(`m["non_existent_key_in_the_map"] ==>`, m["non_existent_key_in_the_map"]) //0
}
