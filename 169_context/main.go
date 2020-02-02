package main

import (
	"context"
	"fmt"
)

func main() {
	definitions()
	fmt.Println("------------------------")
	contextExample1()
	fmt.Println("------------------------")
	contextWithCancelExample()
}

func definitions() {
	fmt.Println(`
	Context helps to close all goroutines related to some process in order to avoid
	leaking in goroutines, that is when goroutines gets running when the process that launched it is not
	running anymore.

	Context can also help ypu with passing variables that are related with a request.
	https://blog.golang.org/context
	`)
}

func contextExample1() {
	ctx := context.Background()
	printContext(ctx)

}

func contextWithCancelExample() {
	ctx := context.Background()
	ctx, _ = context.WithCancel(ctx)
	printContext(ctx)
}

func printContext(ctx context.Context) {
	fmt.Println("Context is: \t", ctx)
	fmt.Println("\nError is", ctx.Err())
	fmt.Printf("Context type is ==> %T\n", ctx)
}
