package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println("Current Time is: ", currTime)
	fmt.Printf("Type of current Time is: %T\n", currTime)

	//Formatted Date
	formattedDate := currTime.Format("02-01-2006")
	fmt.Println("Current Date is: ", formattedDate)

	//Formatted Day
	formattedDay := currTime.Format("Monday")
	fmt.Println("Today's Day is: ", formattedDay)

	// Formatted Time
	formattedTime := currTime.Format("15:04:05")
	fmt.Println("Current time is: ", formattedTime)

	//Todays date with current Day and Time
	date := currTime.Format("02-01-2006, Monday, 15:04:05")
	//date := currTime.Format("02-01-2006, Monday, 03:04 PM") // now the time is in 12 hrs. format
	fmt.Println("Todays date time:", date)

	// Date time parses
	layoutStr := "2006-01-02"
	dateStr := "2023-11-25" // this is in String sormat
	formatTime, _ := time.Parse(layoutStr, dateStr)
	fmt.Println("Formatted Time: ", formatTime)

	//Add 1 more day in currentTime
	newDate := currTime.Add(24 * time.Hour) // 24 ressamble 1 day or 24 hrs.... if re want to increase 3 days then (72 * time.Hour)
	fmt.Println(newDate)
	formattedNewDate := newDate.Format("02-01-2006, Monday, 15:04:05")
	fmt.Println("Todays next day:", formattedNewDate)
}
