package main
import "fmt"

func main(){
	a:=42
	b:=42.56

	fmt.Println("a variable is",a, "its type is")
	fmt.Printf("%T\n",a)
	fmt.Println("b variable is",b,"its type is")	
	fmt.Printf("%T\n",b)
}