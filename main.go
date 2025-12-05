package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"os"
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

	//Decode
	jsonStr := `{"name":"Apip","15"}`
	var person2 Person
	json.Unmarshal([]byte(jsonStr), &person2)
	fmt.Println(person2.Name)

	//Array 
	people := []Person{
		{Name: "Apip", Age: 15},
		{Name: "Budi", Age: 20},
	}
	peopleJson, _ := json.Marshal(people)
	fmt.Println(string(peopleJson))

	//Map
	data := map[string]string{
		"city":   "Jakarta",
		"country": "Indonesia",
	}
	mapJSON, _ := json.Marshal(data)
	fmt.Println(string(mapJSON))

	//Streaming Decoder
	jsonStream := `{"name":"Apip","age":15}{"name":"Budi","age":20}`
	decoder := json.NewDecoder(strings.NewReader(jsonStream))
	var person3 Person
	decoder.Decode(&person3)
	fmt.Println("Streaming Decoder: ", person3)

	//Streaming Encoder
	file, _ := os.Create("output.json")
	encoder := json.NewEncoder(file)
	encoder.Encode(person)
	file.Close()
	fmt.Println("Data written to output.json")
}