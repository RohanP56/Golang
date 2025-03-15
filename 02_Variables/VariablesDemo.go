package main

import "fmt"

func main() {
	//string
	//var name string = "golang"
	var name = "golang"
	fmt.Println(name)

	//number
	var height float32 = 55.3
	fmt.Printf("age is: %.2f\n", height)
	fmt.Printf("Type of age is: %T\n", height)

	//boolean
	var isAdult bool = true
	fmt.Println(isAdult)

	//short hand syntax
	name2 := "Python"

	//printing using format specifier
	fmt.Printf("Language name is %s\n", name2) // \n for goiung nextLine
	fmt.Printf("Language name is %s\n", name)
}
