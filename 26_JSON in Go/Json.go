package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person := Person{Name: "John", Age: 45, IsAdult: true}
	fmt.Println("Person Data:", person)

	//Convert person in to JSON (Encoding / Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error occur while marshelling", err)
		return
	}
	//fmt.Println("Person data is:", jsonData) ---> It's a array of data, so we can't directly access it
	fmt.Println("Person data is:", string(jsonData))

	//Convert person in to JSON (Decoding / Unmarshalling)
	var personData Person //create a variable of type "Person"
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error occur during Decding:", err)
	}
	fmt.Println("Person Data is: ", personData) // it's aready string no need to convert string
}
