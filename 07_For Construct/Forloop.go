package main

import "fmt"

// for -> Only construct in go for looping

//while loop using "for"
func whileLoop(i int) {
	for i <= 5 { // no need to user (), means we don't have to declare condition inside bracket
		fmt.Print(i, " ")
		i++
	}
}

// Implement infinite loop
func infiniteLoop(i int) {
	// in infinite loop we don't have to dec;are any condition
	for {
		fmt.Print(i, " ")
		i++
	}
}

// forloop using for keyword
func forLoop(n int) {
	for i := 0; i <= n; i++ {
		fmt.Print(i, " ")
	}
}

// for loop using range
func forLoopUsingRange(n int) {
	// range starts form 0 and does not include "n"
	for i := range n {
		fmt.Print(i, " ")
	}
}

// function to print all character in side a String
func stringChar() {
	str := "Hello World"
	fmt.Println("Length of given String: ", len(str))
	for i, ch := range str {
		fmt.Printf("At index %d Character is: %c\n", i, ch)
	}
}

func main() {
	/*whileLoop(1)
	infiniteLoop(1)
	forLoop(10)
	forLoopUsingRange(10)*/
	stringChar()

}

/* In go theres only one keyword "for" using for we implements all types of loops break ans continue works as usual*/
