package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	definitions()
	marshallExample1()
}

func definitions() {
	fmt.Println(`
	https://golang.org/ ==> Package Documentation ==> encoding ==> json
	https://godoc.org/encoding/json ==> marshal ==> in that example a struct is taken into a json
	https://godoc.org/encoding/json ==> unmarshaling ==> takes []byte(...{json object}...) ==> into  a desired value
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
	os.Stdout.Write(b)
	return b
}
