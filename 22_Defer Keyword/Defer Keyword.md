# Defer
- Defer keyword use to change the flow of the program.
`defer func(parameters) (returnType) {}`
- If we use **Defer** keyword with any function, that function will execute before end of the progamm.
``` Go
func main() {
	fmt.Println("Programm Started...")
	defer fmt.Println("Middle of the Program....")
	fmt.Println("Programm Ended...")
}

            //Output
`
Programm Started...
Programm Ended...
Middle of the Program....
`
```
- **If we use multiple `defer`, then the execution will happen `LIFO` pattern, When line by line execution happen it will create a stack and if any function has defer that will be added into `stack`... After normal functions execution the stack items will be execute.**
``` Go
func main() {
	sum(5, 8)
	defer sub(5, 8)
	defer mul(5, 8)
	mod(8, 3)
}
            //Output
`
Sum: 13
Mod: 2
Mul: 40
Subs: -3
`
```

### Where to use `difer`?
When we use some file or OS related operation then always we have to close the operations and some times we forgot to close the operations so when we do the operation next of thet we declare the closing operations with `difer` keyword.