# String
- All **String** related methods are inside **strings** package.

### Slice:
- `sliceVariable := strings.Split(stringData, ",")`
- This function returns data in slice format, based on any separator.
``` Go
func main() {
	data := "apple,banana,orange"
	parts := strings.Split(data, ",")
	fmt.Println(parts)
}
```
### Count:
- **This funcion used to count any character or words frequency in a String**
- `integerVariable := strings.Count(stringData, "word or character")`
- This function returns a Integer value based on how many times the data is present
``` Go
func main() {
	str := "agoplddemzaeaoddnddhkjcns"
	countD := strings.Count(str, "d")
	fmt.Println("Total no of d is:", countD)
}
```
### TrimSpace:
- **This function used to trim the starting and ending spaces**
- `stringVariable := strings.TrimSpace(givenString)`
``` Go
func main() {
	str := "   agoplddemzaeaoddnddhkjcns    "
	trimmedString := strings.TrimSpace(str)
	fmt.Println("Trimmed String is:", trimmedString)
}
```
### Join
- **Used to cancateneate multiple String**
- This function takes **array of Strings** and a **Separator**, using this separator separte different Strings.
- `newString := strings.Join(arrayofStrings, Separator)`
``` Go
func main() {
	str1 := "Rohan"
	str2 := "Pramanik"
	fullName := strings.Join([]string{str1, str2}, " ")
	fmt.Println("Full Name:", fullName)
}
```