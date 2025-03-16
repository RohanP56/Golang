package main

import "fmt"

func division(a float64, b float64) (float64, error) { // like error we can also use string as return message
	if b == 0 {
		return 0,
			fmt.Errorf("**Denominator must be >= 1") //Thos is error message
	}

	return (a / b), nil // here no error that's wy we have pass "nil"
}

func main() {
	ans1, _ := division(10, 2) // here we have discard the error section, if we daon't want to use the error section then user ("_")
	fmt.Println("Answer of division is:", ans1)
	ans2, err := division(10, 0)
	fmt.Println("Answer of division is:", ans2, "\nError is:", err)
}
