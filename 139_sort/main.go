package main

import (
	"fmt"
	"sort"
)

func main() {
	definitions()
	example1SortInts()
	example1SortStrings()
}

func definitions() {
	fmt.Println(`
	Sort ===> https://golang.org/pkg/sort/
	`)
}

func example1SortInts() {
	sInts := []int{3, 4, 7, 9, 2}
	fmt.Println("initial slice of ints", sInts)
	sort.Ints(sInts)
	fmt.Println("Final slice after sort", sInts)
}

func example1SortStrings() {
	sStrings := []string{"hello", "this", "is", "a", "new", "try"}
	fmt.Println("initial slice of strings", sStrings)
	sort.Strings(sStrings)
	fmt.Println("Final slice after sorting process", sStrings)
}
