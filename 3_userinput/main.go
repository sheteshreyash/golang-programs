package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "Welcome to the Go programming language!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give rating for our services (1-5):")

	// comma ok or error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("You rated us:", input)
	fmt.Println("Thank you for your feedback!")
	fmt.Printf("Type of the rating is %T", input)
}