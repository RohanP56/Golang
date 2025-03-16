package main

import (
	"fmt"
	"time"
)

// Normal "Switch" case
func getMonth(month int) {

	// in golang switch case no need to wite "break"
	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("Invalid input")
	}
}

// multi control switch
func getDateTime() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("It's Workday")
	}
}

// type switch
func whoAmI(i interface{}) { // interface{}: means this function can take any type if input
	switch i.(type) {
	case int:
		fmt.Println("It's an Integer")
	case string:
		fmt.Println("It's a String")
	case bool:
		fmt.Println("It's a Boolean")
	default:
		fmt.Println("It's undefines")

	}
}

func main() {
	month := 12
	fmt.Print("The month is: ")
	getMonth(month)
	getDateTime() // it will automatically take dateTime from "os"
	whoAmI(5)
	whoAmI(true)
	whoAmI("true")
}
