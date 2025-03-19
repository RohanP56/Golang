package main

import "fmt"

//This counter function does not recieve any thing, but return another function whci return integer.
func counter() func() int {
	var count int = 5

	return func() int {
		count++
		return count
	}
}

func main() {
	increment := counter() // here we recieve a function, because counter returns a function
	fmt.Println(increment())
	fmt.Println(increment())
}
