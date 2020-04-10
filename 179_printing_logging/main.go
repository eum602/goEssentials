package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	example1()
	example3()
	example4()
	example5()
	fmt.Println(example6())
	example2()
}

func definitions() {
	fmt.Println(`
	Defer statements allow us to think about closing each file right after opening it, guaranteeing that, 
	regardless of the number of return statements in the function, the files will be closed.
	`)

}

func example1() {
	fmt.Println("*************Print logs example*************")
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()  //closing at the end
	log.SetOutput(f) //setting that all log information input will go to the log.txt file

	//Now simulating an error to write it to a file
	f2, err := os.Open("non-existent-file.txt")
	if err != nil {
		log.Println("An error has happened: ", err) //this log will be written to log.txt
	}
	defer f2.Close()
}

func example2() {
	fmt.Println("Example 2: log.Fatal == os.Exit(1) and in that case deferreds are not executed!!")
	defer deferred() //this will not be executed because deferred functions will not be executed.
	fmt.Println("*************Log Fatal example: Log fatal invokes os.Exit(1) ==> which means exits with error*************")
	f2, err := os.Open("non-existent-file.txt")
	if err != nil {
		log.Fatal(err) //exiting with os 1 status code ==> exits with error
	}
	defer f2.Close()
}

func example3() {
	fmt.Println("Example 3: Using defer to copy from one file to another")
	f, err := os.Create("log1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()  //closing at the end
	log.SetOutput(f) //setting that all log information input will go to the log.txt file

	written, err := CopyFile("copy", "log.txt")
	if err != nil {
		log.Println("Error in example 3", err)
	}
	log.Println("Successfull copy", written)
}

func example4() {
	fmt.Println("Example 4: printing a last message after return with defer")
	i := 0
	defer fmt.Println(i) //it will print "0" after return returns 1.
	i++
	return
}

func example5() {
	fmt.Println("Example 5: printing in reverse order with defer")
	for i := 0; i < 4; i++ {
		//Deferred function calls are executed in Last In First Out order after the surrounding function returns.
		defer fmt.Print(i, "\n") //3210
	}
}

func example6() (i int) {
	//In this example, a deferred function increments the return value i after the surrounding
	//function returns. Thus, this function returns 2:
	fmt.Println("Example 6: changing the returned message with defer!!")
	defer func() { i++ }()
	return 1
}

func deferred() {
	fmt.Println("Deferred functions will not be executed when os.Exit(1) is invoked")
}

//CopyFile ...
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close() //runs after return is executed. Guarantee even when errors happens the file will be closed.

	return io.Copy(dst, src)
}
