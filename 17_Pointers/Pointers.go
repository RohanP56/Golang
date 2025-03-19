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

func main() {
	num := 1
	/*changeNum(num)
	fmt.Println("After chnageNum in main:", num) //1*/

	changeNum(&num)
	fmt.Println("After chnageNum in main:", num)
}
