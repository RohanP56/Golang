//Slices are same Arraylist in Golang

package main

import (
	"fmt"
	"slices"
)

func main() {
	// declaration
	var nums_1 []int // currently it's nil

	//length of slice
	fmt.Println("Size of nums_1:", len(nums_1))

	// we can also declare a Slice with predefine capacity
	var nums_2 = make([]int, 2)
	fmt.Println("Capacity of nums_2:", cap(nums_2))

	// Adding element in nums
	nums_2 = append(nums_2, 10)

	// Adding element on Particular index
	nums_2[0] = 3

	fmt.Println(nums_2)

	// Copy a Slice
	var nums_3 = make([]int, 0, 5) // declare nums_3
	//Adding value in nums 3
	nums_3 = append(nums_3, 10)
	nums_3 = append(nums_3, 20)
	nums_3 = append(nums_3, 30)
	var nums_4 = make([]int, len(nums_3)) // declare nums_4
	copy(nums_4, nums_3)
	fmt.Println("Nums 3:", nums_3)
	fmt.Println("Nums 4:", nums_4)

	//Slice operator
	var nums_5 = []int{1, 2, 3, 4, 5}
	fmt.Println(nums_5[0:3])

	//check two slices are same or not
	var nums_6 = []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Equal(nums_5, nums_6))

	// 2-D Slices
	var nums_7 = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(nums_7)
}
