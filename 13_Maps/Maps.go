package main

import (
	"fmt"
	"maps"
)

func main() {
	//Creating Map with (make) function
	m := make(map[string]string)

	// Creating map with out make function
	m1 := map[string]int{"price": 100, "phones": 3}
	fmt.Println("Map with-out make function:", m1)

	// Adding element in Map
	m["India"] = "in"
	m["China"] = "ch"
	m["USA"] = "us"

	// getting element
	fmt.Println("India ->", m["India"])
	fmt.Println("USA ->", m["USA"])
	fmt.Println("China ->", m["China"])

	//Printing full map with key and value
	fmt.Println("Map with make function:", m)

	//loop through map
	for index, value := range m {
		fmt.Printf("Key is %s and value is %s \n", index, value)
	}

	//length of map or how may item exists in map
	fmt.Println("Length of Map m:", len(m))

	// Delete any element from map
	delete(m, "India")
	fmt.Println(m)

	//check any key is present on map or not
	val, ok := m["China"] // it returns items value and a boolean value
	fmt.Println(val)      // it will print it's key value
	//_, ok := m["China"]  if we want we can skip it
	if ok {
		fmt.Println("Exists")
	} else {
		fmt.Println("Not Exists")
	}

	//Empty full map
	clear(m)
	fmt.Println(m)

	// Equality Check for map
	m3 := map[string]int{"price": 10000, "Samsung": 3}
	m4 := map[string]int{"price": 40000, "Apple": 3}

	fmt.Println(maps.Equal(m3, m4))
}
