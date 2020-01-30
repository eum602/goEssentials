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
		//after that comes {"james":32, eum602:31 ... } ==> this is the composite literal because we are
		//building values together
		"James":  32,
		"eum602": 31,
	}
	fmt.Println(m)
	fmt.Println(`m["eum602"] ==>`, m["eum602"])                                           //31
	fmt.Println(`m["non_existent_key_in_the_map"] ==>`, m["non_existent_key_in_the_map"]) //0
	v, ok := m["non_existent_key_in_the_map"]
	fmt.Println(`v,ok := m["non_existent_key_in_the_map"] ==>`, v, ok) //0,false
	if v, ok := m["non_existent_key_in_the_map"]; ok {
		fmt.Println("Value exists!") //this wont be printed
	} else {
		fmt.Println("\nThe value associated to the key 'non_existent_key_in_the_map' does not exist because v,ok are:", v, ok)
	}

	m["yess"] = 24
	fmt.Printf(`m["yess"] = 24`) //in this way we add a new value to a map.
	println("\n\nPrinting all key:values of some map", `
	for k, v := range m {
		println(k, ",", v)
	} ==>`)
	for k, v := range m {
		println(k, ",", v)
	}
	fmt.Println(`
	Iterating over a slice is identical than iterating over a map, the difference
	is that slice uses the index whilst maps use the key to iterate over`)
	xi := []int{1, 2, 3, 4, 5} //a slice, created again with composite literals
	for i, v := range xi {
		println(i, ",", v)
	}
}
