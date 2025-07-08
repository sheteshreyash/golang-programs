package main

import "fmt"


func main() {
	fmt.Println("Pointers in Golang")

	var ptr *int // ptr is a pointer to an int, it is initialized to nil
	fmt.Println("Pointer value:", ptr)

	myNumber := 42
	var ptr1 = &myNumber // ptr1 now holds the address of myNumber
	fmt.Println("Pointer value after assignment:", ptr1)
	fmt.Println("Value pointed to by ptr1:", *ptr1) // Dereferencing ptr1 to get the value it points to

	// Changing the value of myNumber using the pointer
	*ptr1 = myNumber + 10
	fmt.Println("Value of myNumber after change:", myNumber)

}