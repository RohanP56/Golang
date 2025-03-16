package main

import "fmt"

// function-1:  without return type and parameter
func greet() {
	fmt.Println("Hello world!")
}

// function2: with parameter but no return type
func callName(name string) {
	fmt.Println("Hi,", name)
}

// function3: with parameter and return type is integer
func sum(a int, b int) int { // 1st take input and then write return type
	return (a + b)
}

// here we can also define the return type with it's variable name
func mul(a int, b int) (result int) {
	result = a * b
	return
}

func main() {
	// function-1
	greet()
	// function-2
	callName("Rohan")
	//function-3
	fmt.Println("Sum of given numbers is:", sum(5, 12))

	data := mul(5, 2)
	fmt.Println("Answer is:", data)
}
