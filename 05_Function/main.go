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

// function4: here we can also define the return type with it's variable name
func mul(a int, b int) (result int) {
	result = a * b
	return
}

// function5: in Golang function can return multiple type of values
func getLanguages() (string, string, bool) {
	return "Golang", "JavaScript", true //Here we also have to maintion the return type same order
}

// Function6: Higher order function, which take function as parameter
func higherOrderFunction1(callBack func(a int) int) {
	callBack(5)
}

//Function7: Which does not recieve anything but return a function
func higherOrderFunction2() func(a int) int {
	return func(a int) int {
		return 5
	}
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

	//function-5
	fmt.Println(getLanguages())
	/*we can also destructure the values
	lang1, lang2, lang3 := getLanguages()
	fmt.Println("1st language in function-5:", lang1)
	fmt.Println("2nd language in function-5:", lang2)
	fmt.Println("3rd language in function-5:", lang3)

	// lets we just want to print 2 values to suppress anoter values use (_)*/
	lang1, _, lang3 := getLanguages()
	fmt.Println("1st language in function-5:", lang1)
	fmt.Println("3rd language in function-5:", lang3)

	//Creating a callBack function and assign it into variable
	fn := func(a int) int {
		return 5
	}

	//Function-6
	higherOrderFunction1(fn)

	// function-7
	f1 := higherOrderFunction2()
	f1(6)
}
