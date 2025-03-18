# Array
- **In golang array follows `0` Based indexing**
- **Array Declaration:** `var array_name [array_size] array_data_type`
``` Go
var arr [5]int
        or
arr := []int{1, 2, 3, 4, 5}
        or
arr := [5]int{1, 2, 3, 4, 5}
```
- **Printing Array:** `fmt.Println(array)`
``` Go
fmt.Println("Array:", arr)
```
- **Array Length:** `len(Array_name)`
``` Go
fmt.Println(len(arr))
```

- ### When we create  a array of particular datatype then, while we don't put any values the array initialize with data types default value like java.
    - **For string type array default value will be space(" ")**
    ``` Go
    var nums_3 [5]int
    fmt.Println("Array (arr):", arr)  // [0 0 0 0 0]
    ```
    - But if we want to see actual values we can use fomat specifier method printing or Quoted String.
    ``` Go
    var nums_4 [5]string
	nums_4[3] = "Akash"
	nums_4[1] = "Rohit"
	fmt.Println("Array: ", nums_4)      // Array:  [ Rohit  Akash ]
    // Quoted String
	fmt.Printf("Array: %q\n", nums_4)  // Array: ["" "Rohit" "" "Akash" ""]
    ```