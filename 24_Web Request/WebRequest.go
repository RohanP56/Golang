package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web Service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("There's error in get response", err)
		return
	}
	defer res.Body.Close()               // Closing 'Responce body' stream.
	fmt.Printf("Type of res is %T", res) //*http.Response

	//read the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error While reading response: ", err)
		return
	}

	fmt.Println("response: ", string(data))
}
