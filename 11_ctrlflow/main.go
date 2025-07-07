package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Control flow statement in golang")

	// 1. if else statements
	loginCount := 21
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount < 20 {
		result = "Intermediate User"
	} else {
		result = "Super User"
	}
	fmt.Println(result)

	// if with short variable declaration
	if n:= 10; n < 20{
		fmt.Println(n)
	} else {
		fmt.Println("n is not less than 20")
	}

	// error handling with if
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Success")
	// }



	// 2. switch statement
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Number:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You rolled a one!")
	case 2:
		fmt.Println("You rolled a two!")
	case 3:
		fmt.Println("You rolled a three!")
	case 4:
		fmt.Println("You rolled a four!")
	case 5:
		fmt.Println("You rolled a five!")
	case 6:
		fmt.Println("You rolled a six, you can roll again!")
		fallthrough
	default:
		fmt.Println("Invalid roll")
	}
}