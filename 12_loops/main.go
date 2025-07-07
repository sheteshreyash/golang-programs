package main

import "fmt"

func main()  {
	fmt.Println("Loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("Days of the week:", days)

	// for loop
	for i := 0; i < len(days); i++ {
		fmt.Println("Day", i+1, ":", days[i])
	}

	// for loop with range
	for i := range days {
		fmt.Println("Day", i+1, ":", days[i])
	}

	// for loop with rougevalue
	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 8 {
			break
		}
		if rogueValue == 3 {
			rogueValue++
			continue
		}
		if rogueValue == 7 {
			goto lco
		}
		fmt.Println("Rogue value is:", rogueValue)
		rogueValue++
	}

	// labels
	lco :
	fmt.Println("This is a label and it hit the rogueValue")



	// 3. While loop
	// Go does not have a while loop, but we can use a for loop to achieve

}