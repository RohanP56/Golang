# Variadic Function
- **It's a type of function, which can take "N" numbers of inputs, but the type of inputs will have to be sametype.**
- Syntax similar like JavaScript, spred Operator but here we called it variadic operator
- same as normal function, but when take input parameters then we have to use variadic operator along with variable and datatype. `function (variable ...dataType) returnType`
- This variadic operator operator internally creates a **slice**, so we iterate over slice and return it's value

``` Go
func sum(nums ...int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}
	return total
}

func main() {
	fmt.Println(sum(5, 12, 68, 75, 0, 8962))
}

```
- We can also take different type of Operators using interface.
`func (params ...interface{})`
``` Go
func sum1(params ...interface{}) {
    //code
}
```
- **We can also pass a full *slice* inside veriadic function. by using variadic operator, when we pass a *arguement* we have to use variadic operator**
``` Go
func sum(nums ...int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}
	return total
}

func main(){
    //creating a slice
	nums := []int{1, 5, 8, 3, 4, 8}
	// we can directly pass a slice
	fmt.Println(sum(nums...))
}
```