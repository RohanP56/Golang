package main

import "fmt"

func main() {
	const k = 100
	fmt.Println(k)

	// mutiple const declaration
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port)
	fmt.Println(host)

	//k = k + 5   invalid
}

/*To assign consttant we have to use (const) keywprd*/
