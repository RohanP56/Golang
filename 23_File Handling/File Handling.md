# File Handling

### File Creation
- Any file related work mainly done by OS, that's the reason we need `os` package.

### File Creation:
- **File Creation done by `os` package**
- `os.Create(fileName with Extention)`
- It returns two things: 
    - File type data
    - Error
- After creating file always free resources.
``` Go
func main() {
	file, err := os.Create("Example.txt")
	//if any error exists
	if err != nil {
		fmt.Println("Error occured while creating file:", err)
		return
	}
	// if file created successfully then free the resources
	defer file.Close()
	fmt.Println("File Created Successfully")
}
```


### File Writing:
- **If we want to add anything in the file, we want `io` package**
- Using `io.WriteString` we can add text in file. It takes two parameters
    - Writer(Where to write)
    - What to write or The text
- This function returns two things 
    - `byte` value [This ressamble, how much byte of date we have entered in the file]
    - Error
``` Go
func main() {
	file, err := os.Create("Example.txt")
	// File Writing
	// File Writing
	fileText := "Hello world 'Rohan' this side"
	byt, eor := io.WriteString(file, fileText)
	if err != nil {
		fmt.Println("Error occur during file writing: ", eor)
		return
	}
	fmt.Println("Total added data:", byt, "Byte")
	fmt.Println("Data added into file Successfully")
}
```


### File Reading:
- There's two way to read file
    - Using **Buffer**.
    - **ReadFile** function from **ioutil** package.

- #### File Reading using Buffer
    - Open the file `os.Open(fileName)`
    - It also returns two things
        - File
        - Error
    - To read data we have create `Buffer: A temporary Storage`
    - `file.Read` used to read file, this method takes **buffer** and returns a **integer** and **error**.
    - Now we loop through buffer and print, items which stored inside buffer.
    ``` Go
    func main() {
	    // File Reading using buffer
	    fileItem, eorr := os.Open("Example.txt")
	    if eorr != nil {
		    fmt.Println("Error occur during file Opening: ", eorr)
		    return
	    }
	    buffer := make([]byte, 1024) // we have created a buffer, byte type and size is 1024

	    //readfile and add into buffer
	    for {
		    n, err := fileItem.Read(buffer)
		    if err == io.EOF { // EOF: End of File
			    break
		    }
		    if err != nil {
			    fmt.Println("Error while file Reading ", err)
			    return
		    }
		    // process the buffer
		    fmt.Println(string(buffer[0:n]))
	    }
    }
    ```

- #### ReadFile Function:
    - **ReadFile** function is inside `os` package.
    - It returns two things
        - Array of byte
        - Error

    ``` Go
    func main() {
        content, errr := os.ReadFile("Example.txt")
	    if errr != nil {
		    fmt.Println("Error while Reading file:", errr)
		    return
	    }
	    fmt.Println(string(content)) // We have to pass this inside a string
    }
    ```
**This method is convenient such scenario, where read entire content of a file into memory. That's why some cases we can see memory related error but the buffer method takes chunks of data and add into memory, so where the file size is large we don't use `ReadFile` Method.**