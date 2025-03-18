//Slices are same Arraylist in Golang

package main

import (
	"fmt"
	"slices"
)

func main() {
	// declaration
	var nums1 []int // currently it's nil

	// Type of slice
	fmt.Printf("Type of nums1 Slice: %T\n", nums1)

	//length of slice
	fmt.Println("Size of nums1:", len(nums1))

	// Adding element on Particular index
	fmt.Println(nums1)

	// we can also declare a Slice with make function
	var nums2 = make([]int, 2, 5)
	fmt.Println("Capacity of nums2:", cap(nums2))
	fmt.Println("Length of nums2:", len(nums2))

	// Adding element in nums this element will add in last of the slice
	nums2 = append(nums2, 10)

	fmt.Println(nums2)

	// Copy a Slice
	var nums3 = make([]int, 0, 5) // declare nums_3
	//Adding value in nums 3
	nums3 = append(nums3, 10)
	nums3 = append(nums3, 20)
	nums3 = append(nums3, 30, 88, 99)
	var nums4 = make([]int, len(nums3)) // declare nums_4
	copy(nums4, nums3)
	fmt.Println("Nums 3:", nums3)
	fmt.Println("Nums 4:", nums4)

	//Slice operator
	var nums5 = []int{1, 2, 3, 4, 5}
	fmt.Println(nums5[0:3])

	//check two slices are same or not
	var nums6 = []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Equal(nums5, nums6))

	// 2-D Slices
	var nums7 = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(nums7)
}
