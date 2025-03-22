package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//File Creation
	file, err := os.Create("Example.txt")
	//if any error exists
	if err != nil {
		fmt.Println("Error occured while creating file:", err)
		return
	}
	// if file created successfully then free the resources
	defer file.Close()
	fmt.Println("File Created Successfully")

	// File Writing
	fileText := "Hello world 'Rohan' this side"
	byt, eor := io.WriteString(file, fileText)
	if eor != nil {
		fmt.Println("Error occur during file writing: ", eor)
		return
	}
	fmt.Println("Total added data:", byt, "Byte")
	fmt.Println("Data added into file Successfully")

	/* File Reading using buffer

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
		if err != nil { // EOF: End of File
			fmt.Println("Error while file Reading ", err)
			return
		}
		// process the buffer
		fmt.Println(string(buffer[0:n]))
	}*/

	content, errr := os.ReadFile("Example.txt")
	if errr != nil {
		fmt.Println("Error while Reading file:", errr)
		return
	}
	fmt.Println(string(content)) // We have to pass this inside a string
}
