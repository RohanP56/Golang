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

	// Array declare and add element in single line
	nums_1 := []int{1, 2, 3}
	//nums_1 := [3]int{1, 2, 3}
	fmt.Println("Array nums: ", nums_1)

	//Two Dimentional Array
	nums_2 := [][]int{{1, 2}, {3, 4}}
	fmt.Println("Two Dimentional Array nums:", nums_2)

	var nums_3 [5]int
	fmt.Println(nums_3) // All values initialize as ""

	var nums_4 [5]string
	nums_4[3] = "Akash"
	nums_4[1] = "Rohit"
	fmt.Println("Array: ", nums_4) //  Array:  [ Rohit  Akash ]
	//Quoted String
	fmt.Printf("Array: %q\n", nums_4) // Array: ["" "Rohit" "" "Akash" ""]
}
