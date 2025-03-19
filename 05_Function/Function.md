# Function
-  function is declare by using **func** keyword.
- The type of parameter wite after, maintion the variable name.
    `variableName typeOfParameter`
- Function is little bit same of like other language, but the **return** type of the function maintioned after defination
    `func functionName (functionParameter) returnType {}`
``` Go
func sum(a int, b int) int {
    return a+b
}
```
- Function call is same as other language
``` Go
func main() {
    result := sum(5, 10)
    fmt.Println(result)
}
```
- If the type of parameters are same, then we maintion the type of parameter single time.
``` Go
func sum(a, b int) int {
    return a+b
}
```
- **In Golang function can return multiple type of values** *Here we also have to maintion the return type same order*
``` Go
func getLanguages() (string, string, string) {
	return "Golang", "JavaScript", "C"
}

//we can call this normal way or after destructure
func main(){
	fmt.Println(getLanguages())

	//we can also destructure the values
	lang1, lang2, lang3 := getLanguages()
	fmt.Println("1st language in function-5:", lang1)
	fmt.Println("2nd language in function-5:", lang2)
	fmt.Println("3rd language in function-5:", lang3)

    // lets we just want to print 2 values to suppress anoter values use (_)*/
	lang1, _, lang3 := getLanguages()
	fmt.Println("1st language in function-5:", lang1)
	fmt.Println("3rd language in function-5:", lang3)
}
```
- here we can also define the return type with it's **variable name**
`functionName (parameters) (ReturnVariableWithParameter) {code}`
``` Go
func mul(a int, b int) (result int) {
	result = a * b
	return
}
```

## First Class Citizen
- **Like JavaScript functions are also *first class citizen* in Golang.**
- Means, we can store function inside a variable.
- We can also pass function as arguement inside another function.
- We can also return a function, from another function.

## Higher Order Function
- **When function take another function as argement.** `func functionName (parameterTypeFunction (paramater) returnType)`
``` Go
func higherOrderFunction(callBack func(a int) int) {
    //code
}
```
``` Go
package main

import "fmt"

// Higher-order function that takes a function as an argument
func applyOperation(a int, b int, operation func(int, int) int) int {
    return operation(a, b)
}

// A simple function to add two numbers
func add(x, y int) int {
    return x + y
}

func main() {
    result := applyOperation(5, 3, add) // Passing 'add' function as an argument
    fmt.Println("Result:", result) // Output: Result: 8
}

```

- **When function returns another function as output.**
`functionName () (returnFunction (parameterOfReturnFunction) typeOfReturn Function)`
``` Go
func higherOrderFunction2 () (func(a int) int) {
	//code
}
```
``` Go
package main

import "fmt"

// Function that accepts a callback function
func processNumbers(a int, b int, callback func(int, int)) {
    callback(a, b) // Executing the callback function
}

// A callback function that prints the sum of two numbers
func printSum(x, y int) {
    fmt.Println("Sum:", x+y)
}

func main() {
    processNumbers(4, 6, printSum) // Passing 'printSum' as a callback
}

```
