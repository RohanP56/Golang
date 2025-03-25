# Web Request
- A `Web Request` referes to an HTTP request made to a web Server. These request are used to retrive or send data over internet.
- In Go, we can make web request using `net/http` package, which provides functions to create and send HTTP requests as well as handle request.
### Get Request:
- Get request done by `http.Get` method.
- This method returns, 2 value 
    - response
    - error
- **The returnded response is `Response` type**
``` Go
func main() {
	fmt.Println("Web Service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("There's error in get response", err)
		return
	}
	defer res.Body.Close()               // Closing 'Responce body' stream.
	fmt.Printf("Type of res is %T", res) //*http.Response
}
```
### Read the response body:
- We read response using `ReadAll` function which lays inside `ioutil`
- This function also returns two types value
    - data
    - err
- `ReadAll` Returns **array of bytes** of data.
- We have to concatenate the data as a String.
``` Go
func main() {
    //read the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error While reading response: ", err)
        return
	}
    fmt.Println("response: ", string(data))
}
```