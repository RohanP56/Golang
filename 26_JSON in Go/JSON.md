# JSON Parsing
- In golang there's apackage called `encoding/json`, which used **encode/decode json**

### Encoding and decoding of JSON in Go:
- **Define a Struct:** Define a struct that represent the JSON data Structure. Each field in the struct have a tag specifying the JSON key associated with it.
- **Marshalling(Encoding):** use `json.Marshal`to convert a Go struct into a JSON-encoded byte array. It exists inside `encoding/json` package.
    - This return two Things
        - Array of byte data
        - Error
``` Go
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person := Person{Name: "John", Age: 45, IsAdult: true}
	fmt.Println("Person Data:", person)

	//Convert person in to JSON (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error occur while marshelling", err)
		return
	}
	//fmt.Println("Person data is:", jsonData) ---> It's a array of data, so we can't directly access it
	fmt.Println("Person data is:", string(jsonData))
}

```
- **Unmarshalling(Decoding):** use `json.Unmarshal` to convert a JSON-encoded byte array into a Go struct. It exists inside `encoding/json` package.
    - This method take two things.
        - JSON Data.
        - Address of the variable, where it would store the *decoded* string.
    - It returns one thing
        - Error.
``` Go
//Convert person in to JSON (Decoding / Unmarshalling)
	var personData Person //create a variable of type "Person"
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error occur during Decding:", err)
	}
	fmt.Println("Person Data is: ", personData) // it's aready string no need to convert string
```