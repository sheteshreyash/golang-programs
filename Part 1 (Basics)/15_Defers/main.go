package main

import (
	"fmt"
)

func main()  {
	// Defer is used to delay the execution of a function until the surrounding function returns
	// It is often used for cleanup actions, such as closing files or releasing resources
	// Deferred functions are executed in LIFO (Last In, First Out) order
	// This means that the last deferred function will be executed first
	// Deferred functions are executed even if a panic occurs in the surrounding function
	// Deferred functions can also be used to ensure that certain actions are always performed, regardless often if an error occurs
	fmt.Println("Defer in golang")
	defer fmt.Println("This is deferred")
	defer fmt.Println("This is also deferred")

	deferExample()
}

func deferExample() {

	for i := 0; i < 5; i++ {
		defer fmt.Println("Deferred in loop:", i)
	}
	fmt.Println("Loop completed, deferred functions will now execute")

}