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
- **4th** Post Data using `http.post()` method.
	- **http.post():** This method takes three parameters, 
		- **URL String:** In which `URL` we want to send our data.
		- **Content Type:** The type of Content which we want to send. we write this as `"application/json"`
		- **Reader:** *The body we parse is a type of `Reader`*
	- The Sequence of passing Parameter will be same as maintioned.
- **5th** The `http.post()` returns 2 things
	- Response
	- Error
- **6th** After getting response, close the resources.
- **7th** Now, coonvert the response in redable format.
	- We done this using `ioutil.ReadAll()` method.

``` Go
// POST Request
func postRequest() {
	// URL for Posting Data
	myURL := "https://jsonplaceholder.typicode.com/todos/post"

	// Creating Data
	todo := Todo{
		UserID:    10,
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

	//Convert String into Reader
	JsonReader := strings.NewReader(jsonString) // data is ready to post

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
	fmt.Println("Response Data:", string(data))  // response data
}
```
