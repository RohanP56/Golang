# Data Conversion
- `fmt.Printf("Type of num is: %T\n", num)` Using this we can know the datatype.
``` Go
func main() {
	var num int = 45
	fmt.Printf("Type of num is: %T\n", num)
}
```
### Convert Integer to Float:
- `var variableName DataTypeWantBeChange = DataTypeWantBeChange(oldVariable)`
``` Go
func main() {
	var num int = 45
	var floatNum float64 = float64(num)
	fmt.Println("Value of floatNum:", floatNum)
	fmt.Printf("Type of floatNum: %T\n", floatNum)
}
```
### Convert Integer to String:
- `declareStringVariable := strconv.Itoa(integerVariable)`
- **strconv** Package used to convert into String.
- **Itoa** means [Integer to a(a is 1st chacater that's why denoted by a)]
``` Go
func main() {
	var num int = 45
	str := strconv.Itoa(num)
	fmt.Println("Value of num:", str)
	fmt.Printf("Type of num is: %T\n", str)
}
```
### Convert String to Integer:
- `declareIntegerVariable, error := strconv.Atoi(stringVariable)`
- **Atoi** means [String to Integer]
- This function also returns a **error**, which we can supress.
``` Go
func main() {
	numberString := "1234"
	numberInteger, _ := strconv.Atoi(numberString)
	fmt.Println("Integer of numberInteger:", numberInteger)
	fmt.Printf("Type of num is: %T\n", numberInteger)
}
```
### Convert String to Float:
- `declareFloatVariable, error := strconv.ParseFloat(stringVariable, bitValue)`
- With **strconv** function, we have to use **ParseFloat** .. which parse into float value.
``` Go
func main() {
	numStr := "3.14"
	numFlt, _ := strconv.ParseFloat(numStr, 64)
	fmt.Println("float of numFlt:", numFlt)
	fmt.Printf("Type of numFlt is: %T\n", numFlt)
}
```