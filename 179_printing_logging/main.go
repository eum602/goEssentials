package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	example1()
}

func example1() {
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
