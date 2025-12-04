package main

import (
	"encoding/json"
	"fmt"
	// "log"
	// "strings"
)


type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
}

func main() {
	//Encode
	person := Person{Name: "Apip", Age: 15}
	jsonBytes, _ := json.Marshal(person)
	fmt.Println(string(jsonBytes))
}