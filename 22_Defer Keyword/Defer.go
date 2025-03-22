package main

import "fmt"

func sum(a int, b int) {
	fmt.Println("Sum:", a+b)
}
func sub(a int, b int) {
	fmt.Println("Subs:", a-b)
}
func mul(a int, b int) {
	fmt.Println("Mul:", a*b)
}
func mod(a int, b int) {
	fmt.Println("Mod:", a%b)
}

func main() {
	sum(5, 8)
	defer sub(5, 8)
	defer mul(5, 8)
	mod(8, 3)
}
