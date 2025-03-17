package main

import "fmt"

func main() {
	//Array declaration
	var arr [5]int

	//Array length
	fmt.Println("Length of Array is:", len(arr))

	//Adding element in array
	for i := range len(arr) {
		arr[i] = i + 1
	}

	//printing array
	fmt.Println("Array (arr):", arr)

	// Array declare and add element in singlr line
	nums_1 := [3]int{1, 2, 3}
	fmt.Println("Array nums:", nums_1)

	//Two Dimentional Array
	nums_2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("Two Dimentional Array nums:", nums_2)
}
