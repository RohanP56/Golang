package main

import "fmt"

//range use to iterate over Data Structure
func main() {
	// Iterate through Slice
	nums := []int{1, 2, 3, 6, 8, 9}
	sum := 0

	for i := range len(nums) {
		sum += nums[i]
		fmt.Println("Index", i, "-> Value", nums[i])
	}
	fmt.Println()
	fmt.Print("Total sum:", sum)

	// Iterate through Map
	m := map[string]int{"in": 91, "us": 1, "ch": 86, "bd": 880}
	//This only return key
	for key := range m {
		fmt.Println(key)
	}
	// This returns key, value both
	for key, val := range m {
		fmt.Println("Country", key, "->", "Code", val)
	}

	//Iterate throuh String
	for i, c := range "golang" {
		//fmt.Println(i, c) // return index and unicode point rune
		fmt.Println(i, "->", string(c)) // return index and string
	}
}
