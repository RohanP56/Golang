# Slices

**Slices are same Arraylist in Golang**

- ### How to declare a slice:
*Declaration same as Array but we don't declare the size.*
```Golang
var nums []int  // currently it's nil
        or
var nums := []int{1, 2, 3, 4, 5}
        or
nums := []int{}
```
- ### We can also create a slice with length and capacity:
    - *Make function takes 2 parameter (length and capacity)*
    - When length surpassed the capacity capacity will became double of given capacity. 
```Go
var nums_2 = make([]int, 2, 5)  // 1st one for length and 2nd one capacity
```

- **Check type of a Slice**
``` Go
fmt.Printf("Type of nums_1 Slice: %T\n", nums)
```

## Method Of Silce

- **Length of Slice:** `len(nums)` same as Array
- **Capacity of Slice:** `cap(nums_2)` Maximum number of elements can be fit
- **Adding element in Slice:** `nums_2 = append(nums_2, 10)` element will be added into last of the Slice.
- **Adding element in Particular index:** `slice[index] = value`
- **Adding element on a Particular Position:** `nums_2[index] = value`
- **Copying element from one slice to another:** `copy(src, dest)`

```Golang
// Copy a Slice
var nums_3 = make([]int, 0, 5) // declare nums_3
//Adding value in nums 3
nums_3 = append(nums_3, 10)
nums_3 = append(nums_3, 20)
nums_3 = append(nums_3, 30, 45, 66)
var nums_4 = make([]int, len(nums_3)) // declare nums_4
copy(nums_4, nums_3)
fmt.Println("Nums 3:", nums_3)
fmt.Println("Nums 4:", nums_4)
```
- **Check Given two slice same or Not:** slices.Equal(slice1, slice2)
- **Slice Operator:** `nums_5[S_Index:E_Index]` *return value form Starting index to (Ending index-1), ending index is not included.* 
    - **nums_5[ : E_Index] :** Here only Ending index maintion so Starting index will be 0 automatically. 
    - **nums_5[S_Index : ] :** Here only Ending index maintion so Starting index will be 0 automatically. 

#### 2-D Slices
``` Go
var nums_7 = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
fmt.Println(nums_7)
```