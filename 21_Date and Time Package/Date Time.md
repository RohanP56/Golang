# Date and Time

- in Golang Date and Time Package, the reference date and time for formatting are fixed as **2006-01-02 15:04:05** *January 2, 2006, 3:04:05 PM, MST (Mountain Standard Time)*
- In goland there's a specific Datatytpe to Store Time **time.Time**
``` Golang
func main() {
	currTime := time.Now()
	fmt.Println("Current Time is: ", currTime)
	fmt.Printf("Type of current Time is: %T\n", currTime)
}
```

### Formatted Date
- To get Formatted Date we can't use "dd-mm-yyyy", this format
- we have to pass "02-01-2006"
- `newDate := currTime.Format("02-01-2006") `
``` Golang
func main() {
	currTime := time.Now()
	//Formatted Date
	formattedDate := currTime.Format("02-01-2006")
	fmt.Println("Current Date is: ", formattedDate)
}
```

### Formatted Day
- Here we have to pass **Monday**
- `newDay := currTime.Format("Monday")`
``` Golang
func main() {
	currTime := time.Now()
	//Formatted Day
	formattedDay := currTime.Format("Monday")
	fmt.Println("Today's Day is: ", formattedDay)
}
```

### Formatted Time
- Here we have to pass **15:04:05**
``` Golang
func main() {
	currTime := time.Now()
    // Formatted Time
	formattedTime := currTime.Format("15:04:05")
	fmt.Println("Current time is: ", formattedTime)
}
```
- Using this we can get all to geather
``` Go
func main() {
	currTime := time.Now()
    //Todays date with current Day and Time
	date := currTime.Format("02-01-2006, Monday, 15:04:05")
	//date := currTime.Format("02-01-2006, Monday, 03:04 PM") // now the time is in 12 hrs. format
	fmt.Println("Todays date time:", date)
}
```
### Date and Time Parse
- This method parse the given date time which is in String format.
- This function take given date time and a Layout String, in layout String we define the go's default date, time and store it in given date date, time format.
- if date and time format is different then we have to provide the "layOutString" in different Format.
    - given Date: **(21/03/2025)**
    - LayoutString: **(02/01/2006)**
``` Go
func main() {
	// Date time parse
	layoutStr := "2006-01-02"
	dateStr := "2023-11-25" // this is in String sormat
	formatTime, _ := time.Parse(layoutStr, dateStr)
	fmt.Println("Formatted Time: ", formatTime)
}
```

### Add 1 day with Today:
- To calculate this we use **Add** function, with currentTime
- inside this function, we pass how many time we want to increase in **hour** format and multiply with **time.Hour**
- Then format the value

``` Go
func main() {
    currTime := time.Now()
    //Add 1 more day in currentTime
	newDate := currTime.Add(24 * time.Hour) // 24 ressamble 1 day or 24 hrs.... if re want to increase 3 days then (72 * time.Hour)
	fmt.Println(newDate)
	formattedNewDate := newDate.Format("02-01-2006, Monday, 15:04:05")
	fmt.Println("Todays next day:", formattedNewDate)
}
```