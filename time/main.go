package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time is: ", currentTime)
	fmt.Println("Current time with format: ", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Current time with format: ", currentTime.Format("02-01-2006 03:04:05 PM Monday"))

	layoutStr := "2006-01-02"
	dateStr := "2024-04-20"
	formattedTime, _ := time.Parse(layoutStr, dateStr)
	fmt.Println("Formatted time: ", formattedTime)

	// add 1 day in current time
	oneDay := currentTime.Add(24*time.Hour)
	fmt.Println("One day after current time: ", oneDay)
	formatted_new_date := oneDay.Format("2006-01-02 Monday")
	fmt.Println("One day after current time with format: ", formatted_new_date)
}
