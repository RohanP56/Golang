# CRUD in GO

### GET Method:

- Insted of getting jSON data traditional way, we will create a `Structure` and inside that structure, we pass JSON data and parse the data as we wish.

```Go
// Older data
func main() {
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
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while Reading:", err)
		return
	}
	fmt.Println("Data: ", string(data)) // get json data
}
```

- The `todo Structure` field will have to be same as `returned JSON` data. either inside JSON initializer the name should be same or the variable name for JSON initializer would be same, id both are different then we got error.

```Go
// Structure to format JSON data
type Todo struct {
	UserID    int    `json: "userId"`
	Id        int    `json: "id"`
	Title     string `json: "title"`
	//title     string `json: "Title"`  ----> This is right, coz in variable name the name is same as recieved JSON.
	Completed bool   `json: "completed"`
	//Comp bool   `json: "Completed"`  ----> This is wrong cause, both area we have change the variable name.
}
```

```json
// Returned JSON Data
{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
```

- **NewDecoder():** New Decoder, decode all data in normal objects and save into todo variable.
  - This function only return **err**

```Go
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
```

**Using this method we can access, our required data... API sends too many data but we don't need all data that's why we use this method.**

### POST Method:

- **1st**, We create the data which we want to post, using the Structure.
- **2nd** Convert the Created Data into `JSON` using **Mahshal** Mathod.
- **3rd** We can't send data directly in `JSON`, so we convert the data into **String**.
- **4th** `http.post()` method expects a **Reader interface**, that's why we have to convert it into a reader using `strings.NewReader()` method.
- **5th** Post Data using `http.post()` method.
  - **http.post():** This method takes three parameters,
    - **URL String:** In which `URL` we want to send our data.
    - **Content Type:** The type of Content which we want to send. we write this as `"application/json"`
    - **Reader:** _The body we parse is a type of `Reader`_
  - The Sequence of passing Parameter will be same as maintioned.
- **6th** The `http.post()` returns 2 things
  - Response
  - Error
- **7th** After getting response, close the resources.
- **8th** Now, coonvert the response in redable format.
  - We done this using `ioutil.ReadAll()` method.

```Go
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
```

### UPDATE Method:

**_There's no function called `http.Put()` in Golang_**

- **1st**, We create the data which we want to update, using the Structure.
- **2nd** Convert data into JSON, using `Marshal`
- **3rd** Convert JSON into String
- **4th** Convert String into reader
- **5th** Create PUT Request, using `http.NewRequest()` method. This method takes 3 parameters.
  - **Method:** http.MethodPut or "PUT"
  - **URL** URL where you want to send data.
  - **Body** Reader data which we have created from the JSON data.
  - This Method also returns 2 things
    - **Request**
    - **Error**
- **6th** Set Request content type `req.Header.Set("Content-type", "application/json")`
- **7th** Create client `client := http.Client{}`
- **8th** Send the request `client.Do(req)`, this method returns two things
  - Response
  - Error
- **9th** now, after getting response body, close response.
- **10th** Convert the response in redable format

```Go
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
```

### DELETE Method:

- We don't send any data to perform `delete request`
- Delete Method doesn't return anything, it is done or not. we can understand by status code.
- Go does not provies `http.Delete()`, we have to make it using `http.NewRequest() method.`
- **1st** Define URL.
- **2nd** Create Delete Request using `http.NewRequest(http.MethodDelete, URL)`, This method returns 2 elements.
  - Request
  - Error
- **3rd** Create Clent.
- **4th** Send request to client
- **5th** Close the response body.
- **6th** Convert the

```Go
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
```
