# Pointers
- Pointer save address of variable or save memory location of variables.
- In golang when we pass any arguement to a function, then the arguement by value copy in parameter.
- Means inside function, we get a copy of the arguement.
- If we change the copy of value inside the function, source value remains same, means the value we pass it's unchnagable.
``` Go
//  By Value
func changeNum(num int) { // here we will get the copy of arguement
	num = 5
	fmt.Println("In ChangeNum: ", num) //5
}

func main() {
	num := 1
	changeNum(num)

	fmt.Println("After chnageNum in main:", num) //1
}
```
- **But some times we need to change yhe argements value**
- *Here's pointer come into picture: we don't pass the number as arugeument, we pass address of that variable*
- **To see any variable *Memory* location use `&`**
`fmt.Println("Memory location of num:", &num)`
- To recieve any pointer or adderess as a Parameter then we have to use **(*)**
- To send any value as argement inside pointer we have to use **(*)**
`func num (a *int)`
``` Go
// By Reference
func changeNum(num *int) {
	*num = 5 // Means, add the 5 value in num location
	fmt.Println("In changeNum", *num)
}

func main() {
    num := 1
    changeNum(&num)
	fmt.Println("After chnageNum in main:", num)
}
```