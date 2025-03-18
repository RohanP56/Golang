# Switch Condition
- Switch condition same as other language
``` Go
switch VariableNameWeWantToSwitch
case 1:
    // do execution
case 2:
    // do execution
default:
    // do execution
```
- `variable.(type)` returns what is the type of given variable
``` Go
int n = 5
fmt.Print(n.(type)) // int
```
- **interface{}:** It indicats prameter, if any function take *interface{}* as a parameter means this function can take any data type as a input.
``` Go
func whoAmI(i interface{}) { // interface{}: means this function can take any type if input
	switch i.(type) {
	case int:
		fmt.Println("It's an Integer")
	case string:
		fmt.Println("It's a String")
	case bool:
		fmt.Println("It's a Boolean")
	default:
		fmt.Println("It's undefines")
	}
}
```