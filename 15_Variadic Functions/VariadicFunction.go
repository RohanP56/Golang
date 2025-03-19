package main

import "fmt"

// we are getting multiple parameters same type
func sum(nums ...int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}
	return total
}

// here we are getting multiple values parameters types
func sum1(nums ...interface{}) {
	fmt.Println(nums)
}

func main() {
	fmt.Println(sum(5, 12, 68, 75, 0, 8962))

	//creating a slice
	nums := []int{1, 5, 8, 3, 4, 8}
	// we can directly pass a slice
	fmt.Println(sum(nums...))
}
