package main

import "fmt"

/*  By Value
func changeNum(num int) { // here we will get the copy of arguement
	num = 5
	fmt.Println("In ChangeNum: ", num) //5
}*/

// By Reference
func changeNum(num *int) {
	*num = 5 // Means, add the 5 value in num location
	fmt.Println("In changeNum", *num)
}

// modify value
func modifyValueByReference(num *int) {
	*num = *num + 20
}

func main() {
	var n int
	n = 10

	var ptr *int // create a pointer (int) defines the data wkich store in "ptr" is integer type
	ptr = &n     // ptr = address of "n"

	// [address of n == ptr] and [value of n = value inside ptr address]
	fmt.Println("The value of n is:", n)
	fmt.Println("The value of n is:", *ptr)
	fmt.Println("The address of n is:", &n)
	fmt.Println("The address of n is:", ptr)

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Updated value: ", value)

	num := 1
	/*changeNum(num)
	fmt.Println("After chnageNum in main:", num) //1*/

	changeNum(&num)
	fmt.Println("After chnageNum in main:", num)
}
