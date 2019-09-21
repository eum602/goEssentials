package main

import "fmt"

func main() {
	//KEYWORDS: reserved words like "var,switch", etc
	//OPERATOR: "+,-,*" , etc
	//OPERANDS:"5+3" => 5 and 3 are operands
	//all specs are shown here:
	//https://golang.org/ref/spec

	//short declaration
	//STATEMENT: each operation, declaration. constant and so forth is a statement, you can have more than one statement in one line, but it is not usual; each statement
	//separated by a semicolon
	a := 42 //here we are declaring
	fmt.Println("a", a)
	//reassigning a value
	a = 42
	fmt.Println("a is reassigned", a)

	//EXPRESSION : joins more than one explicit values, constant values, operators and functions in order to produce another value
	y := 100 + 24 // 100 + 24 is the statement and "y" is the result after evaluation => all of this is a Expression
	//NOTE: a statement is often made up of expressions
	fmt.Println("y declared as a ", y)
}
