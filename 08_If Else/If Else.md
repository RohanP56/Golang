# If Eles Statement
- If else same as other language.
- Here we don't have to use braces when declare condition. `age >= 18`
``` Go
if age >= 18 {
    fmt.Println("Person is: Adult")
}
```
- **If have to start the curly braces in that line where the if ends.**
``` Go
            //We have to use '{' in that line where if exists
if age >= 18 {

}
        //Can't use
if age >= 18 
{
    //This is wong
}
```
-  **Same follows for *else* or *else if*, these will be written where the, if statement's ending curly braces exists**
``` Go
            //We have to use 'else' where 'if' ends
if age >= 18 {

} else
        //Can't use
if age >= 18 
{
    //This is wrong
}
else
```
- If we want to use mutiple **logical** operator then add first two in a braces and calculate onther one with that.
`if logic3 && (logic1 || logic2)` this case go would not remove the braces.
``` Go
if x < 5 && (y > 5 || z < 43>) {
    // Do execution
}
```
