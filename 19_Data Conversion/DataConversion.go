package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 45
	fmt.Println("Integer of num:", num)
	fmt.Printf("Type of num is: %T\n", num)

	// convert the int to float
	var floatNum float64 = float64(num)
	fmt.Println("Float of floatNum:", floatNum)
	fmt.Printf("Type of floatNum: %T\n", floatNum)

	// integer to string
	str := strconv.Itoa(num)
	fmt.Println("String of num:", str)
	fmt.Printf("Type of num is: %T\n", str)

	//String to Integer
	numberString := "1234"
	numberInteger, _ := strconv.Atoi(numberString)
	fmt.Println("Integer of numberInteger:", numberInteger)
	fmt.Printf("Type of num is: %T\n", numberInteger)

	//String to Float
	numStr := "3.14"
	numFlt, _ := strconv.ParseFloat(numStr, 64)
	fmt.Println("float of numFlt:", numFlt)
	fmt.Printf("Type of numFlt is: %T\n", numFlt)
}
