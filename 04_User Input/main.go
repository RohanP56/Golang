package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		var name string
		fmt.Println("Enter Your Name:")
		fmt.Scan(&name) (scan) only read before space (" ")
		fmt.Println("Entered name is:", name)
	*/
	fmt.Print("Enter Your Name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') //redaer will read until a new line
	fmt.Println("Entered name is: Mr.", name)
}
