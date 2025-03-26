package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// using this structure we will format the JSON
type Todo struct {
	UserID    int    `json: "userId"`
	Id        int    `json: "id"`
	Title     string `json: "title"`
	Completed bool   `json: "completed"`
}

// GET Request
func getRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/5")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer res.Body.Close()

	//if response status code not success we can give error
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response:", res.Status)
		return
	}

	/*
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error while Reading:", err)
			return
		}

		fmt.Println("Data: ", string(data)) // get json data
	*/

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding: ", err)
		return
	}
	fmt.Println("Todo:", todo)
	//Here we can access all the data, separately
	fmt.Println("Todo Id:", todo.Id)
	fmt.Println("Todo Title:", todo.Title)
}

// POST Request
func postRequest() {
	// URL for Posting Data
	myURL := "https://jsonplaceholder.typicode.com/todos"

	// Creating Data
	todo := Todo{
		UserID:    10,
		Id:        5,
		Title:     "the Social Network",
		Completed: true,
	}

	// Convert todo Structure in JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in Marshalling:", err)
		return
	}

	// Convert JSON in String
	jsonString := string(jsonData)

	//Convert String into Reader because, (http.Post()) Metod expects a reader interface.
	JsonReader := strings.NewReader(jsonString)

	// Posting the data
	res, err := http.Post(myURL, "application/json", JsonReader)
	if err != nil {
		fmt.Println("Error in Sending resquest:", err)
		return
	}
	defer res.Body.Close() // after getting response, close the resource

	// After posting data Response Status
	fmt.Println("Response Status:", res.Status)

	//convert the response in redable format
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while converting the response in redable format:", err)
		return
	}
	fmt.Println("Response Data:", string(data)) // response data
}

// UPDATE Request
func updateRequest() {
	// URL for Posting Data
	myURL := "https://jsonplaceholder.typicode.com/todos/5"

	// Creating Data
	todo := Todo{
		UserID:    10,
		Id:        5,
		Title:     "Top Gun",
		Completed: false,
	}

	// Convert todo Structure in JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in Marshalling:", err)
		return
	}

	// Convert JSON in String
	jsonString := string(jsonData)

	//Convert String into Reader because, (http.Post()) Metod expects a reader interface.
	jsonReader := strings.NewReader(jsonString)

	// Create PUT Request
	//http.NewRequest(http.MethodPut, myURL, jsonReader)
	req, err := http.NewRequest("PUT", myURL, jsonReader)
	if err != nil {
		fmt.Println("Error occur while put requesting:", err)
		return
	}
	req.Header.Set("Content-type", "application/json") // Request is ready

	//create Client
	client := http.Client{}

	//send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	// close response Body
	defer res.Body.Close()

	// After posting data Response Status
	fmt.Println("Response Status:", res.Status)

	//convert the response in redable format
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while converting the response in redable format:", err)
		return
	}
	fmt.Println("Response Data:", string(data)) // response data

}

// DELETE Request
func deleteRequest() {
	// URL for Deleting Data
	myURL := "https://jsonplaceholder.typicode.com/todos/3"

	//Create Delete request
	req, err := http.NewRequest(http.MethodDelete, myURL, nil) // here we don't have to send any data so "nil"
	if err != nil {
		fmt.Println("Error creating DELETE Request: ", err)
	}

	//create Client
	client := http.Client{}

	//send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	// close response Body
	defer res.Body.Close()

	// After posting data Response Status
	fmt.Println("Response Status:", res.Status)
}

func main() {
	//getRequest()
	//postRequest()
	//updateRequest()
	deleteRequest()
}
