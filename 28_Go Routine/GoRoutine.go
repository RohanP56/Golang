package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello World")
}

func greet(name string) {
	fmt.Println("Hi,", name)            // This line will be printed
	time.Sleep(2000 * time.Millisecond) // wait for 2 second then next work will be executed
	fmt.Println("Hello,", name)         // After 2 second of wait this line will be printed
}

func sum(a, b int) {
	fmt.Println("Sum:", (a + b))
}

/*
	func sub(a, b int) {
		fmt.Println("Subs:", (a - b))
	}

	func mul(a, b int) {
		fmt.Println("Mul:", (a * b))
	}

	func div(a, b int) {
		fmt.Println("Div:", (a / b))
	}
*/
func main() {
	fmt.Println("Inside Main function")
	hello()
	go greet("Rohan")
	sum(8, 2)
	/*
		sub(8, 2)
		mul(8, 2)
		div(8, 2)
	*/
	time.Sleep(1000 * time.Millisecond) // wait ing for 1 second before end the program
}
