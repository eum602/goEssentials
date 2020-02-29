package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	definitions()
	marshallExample1()
	unMarshallExample1()
	unMarshallExample2()
}

func definitions() {
	fmt.Println(`
	https://golang.org/ ==> Package Documentation ==> encoding ==> json
	https://godoc.org/encoding/json ==> marshal ==> in that example a struct is taken into a json
	https://godoc.org/encoding/json ==> unmarshaling ==> takes []byte(...{json object}...) ==> into  a desired value
	https://mholt.github.io/json-to-go/ ==> resource to autogenerate go code from json
	`)
}

func createJSON() interface{} {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Blues",
		Colors: []string{"LowBlue", "DarkBlue", "LightBlue"},
	}

	return group
}

func marshallExample1() []byte {
	group := createJSON()
	b, err := json.Marshal(group) //receives an INTERFACE(so any) and returns a SLICE OF BYTES in "b" variable
	if err != nil {
		fmt.Println("Err", err)
	}
	fmt.Println("\n\n****Unmarshall Examples***")
	os.Stdout.Write(b)
	return b
}

func unMarshallExample1() {
	b := marshallExample1()
	fmt.Println(b)
}

func unMarshallExample2() {
	var myJSON = []byte(`[
		{"Name": "Erick" , "NickName": "eum602", "Tasks": "develop" },
		{"Name": "Alice" , "NickName": "al", "Tasks": "assist" }
	]`) //MyJSON is a SLICE OF BYTES(uint8)

	type Profile struct {
		Name     string
		NickName string
		Tasks    string //if you declare with lowercase the variable is not shown
	}

	var profiles []Profile //also it could have been profiles := []Profile{}

	err := json.Unmarshal(myJSON, &profiles)

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Printf("\n\n%v", profiles)
}
