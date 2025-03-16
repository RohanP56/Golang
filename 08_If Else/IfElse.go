package main

import "fmt"

func checkAge(age int) {
	if age >= 18 {
		fmt.Println("Person is: Adult")
	} else if age <= 18 { // we have to write else if or else, where "if" has ended, if we wite in nextLine it will throws an error
		fmt.Println("Person is: Under Age")
	} else {
		fmt.Println("Person is: Senior")
	}
}

func main() {
	age := 17
	checkAge(age)
}
