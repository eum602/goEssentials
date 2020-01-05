package main

import (
	"fmt"
)

func main() {
	definitions()
}
func definitions() {
	fmt.Println(`==================== Concurrency Vs Parallelism ====================
	2006 => intel release dual core processor
	2007 => Google started to create Golang => Go was created to take advantage of multiple cores.
			Go added concurrecny in a way to be used easily, so go is a programming language to natively take advantage
			 of multiple cores. Concurrency is the hard of the design and the evolution of the go programming language.
	2012 => Go v1 was released

	If you have only one cpu and you apply concurrency then sequence of tassk will be performed one  right after the other.
	If you have more than one cpu, then you will have the opportunity to run code in parallel.

	Concurrency: It is a design pattern. You can  write concurrent code, which has the potential to run in parallel only 
	if you have more than one core tneh that concurrent code can run in parallel. Concurrency does not guarantee that
	code will run in parallel. (https://www.youtube.com/watch?v=oV9rvDllKEg)	
	`)

}
