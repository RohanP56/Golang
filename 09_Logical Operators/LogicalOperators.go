package main

import "fmt"

func cinemaTicket(role string, age int) {
	if role == "student" || age < 18 {
		fmt.Println("Ticket not available")
	} else if role == "majdoor" && age >= 20 {
		fmt.Println("Ticket available")
	} else {
		fmt.Println("House full")
	}
}

func main() {
	cinemaTicket("majdoor", 20)

	// we can declare a veriable inside if constructor
	if age := 5; age >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is a child")
	}
}

// go does not have ternary operator
