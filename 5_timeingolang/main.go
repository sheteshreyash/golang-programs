package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Time Study in Golang")

	// Getting current time
	presentTime := time.Now()
	fmt.Println("Current Time: ", presentTime)

	// Formatting time
	fmt.Println("Current time : ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Creating a specific time and formatting
	createdTime := time.Date(2023, time.October, 3, 13, 0, 0, 0, time.UTC)
	fmt.Println("Created Time: ", createdTime)
	fmt.Println("Created time : ", createdTime.Format("01-02-2006 15:04:05 Monday"))
	
}