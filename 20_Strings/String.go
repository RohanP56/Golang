package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,banana,orange"

	//Split function returns a slice, with comma separated value
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	//Count function used to count any string or character frequency.
	str := "                agoplddemzaeaoddnddhkjcns   "
	countD := strings.Count(str, "d")
	fmt.Println("Total no of d is:", countD)

	// TrimSpace used to trim the starting and ending spaces
	trimmedString := strings.TrimSpace(str)
	fmt.Println("Trimmed String is:", trimmedString)

	//Join used to cancatenate String
	str1 := "Rohan"
	str2 := "Pramanik"
	fullName := strings.Join([]string{str1, str2}, " ")
	fmt.Println("Full Name:", fullName)
}
