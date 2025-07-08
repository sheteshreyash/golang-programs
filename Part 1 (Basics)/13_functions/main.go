package main

import "fmt"

func main()  {
	fmt.Println("Functions in golang")
	greeting()

	result := add(5, 10)
	fmt.Println("The result of addition is:", result)

	proResult := proAdder(5, 10, 15, 30)
	fmt.Println("The result of proAdder is:", proResult)
}

func greeting() {
	fmt.Println("Namaste, Shreyash Ayaan!")
}

func add(a int, b int) int {
	sum := a + b
	return sum
}

func proAdder(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
