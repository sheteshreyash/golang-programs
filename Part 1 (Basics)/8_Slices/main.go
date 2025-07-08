package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Slices in golang")

	var fruitList = []string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// Inserting elements into a slice
	fruitList = append(fruitList, "Elderberry", "Fig", "Grape")
	fmt.Println("Updated fruit list:", fruitList)

	// Slicing a slice
	citrusList := fruitList[1:] // Slicing from index 1
	fmt.Println("Citrus list:", citrusList)

	// slicing with a range
	berryList := fruitList[1:4] // Slicing from index 1 to index 4 (exclusive)
	fmt.Println("Berry list:", berryList)

	// Length and capacity of a slice
	fmt.Println("Length of fruitList:", len(fruitList))
	fmt.Println("Capacity of fruitList:", cap(fruitList))

	// using make to create a slice
	numbers := make([]int, 5) // Create a slice of integers with length 5
	fmt.Println("Numbers slice:", numbers)
	numbers[0] = 1001
	numbers[1] = 4008
	numbers[2] = 300
	numbers[3] = 200
	numbers[4] = 5000
	numbers = append(numbers, 40, 50) // Appending more elements
	fmt.Println("Updated Numbers slice:", numbers)

	// sorting a slice
	sort.Ints(numbers) // Sorting the numbers slice
	fmt.Println("Sorted Numbers slice:", numbers)
	fmt.Println("Are numbers sorted?", sort.IntsAreSorted(numbers))

	// Removing a value from a slice based on index
	var courses = []string{"Go", "Python", "Java", "C++", "JavaScript"}
	fmt.Println("Courses before removal:", courses)
	var index int = 1 // Index of the course to remove
	courses = append(courses[:index], courses[index+1:]...) // Remove the course
	fmt.Println("Courses after removal:", courses)


}