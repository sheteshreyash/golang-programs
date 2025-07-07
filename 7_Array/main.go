package main

import "fmt"

func main()  {
	fmt.Println("Array in golang")

	// Declare and initialize an array
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Cherry"
	fruitList[3] = "jackfruit"

	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Length of Fruit List:", len(fruitList))

	// Declare and initialize an array with values
	vegetableList := [4]string{"Carrot", "Potato", "Tomato", "Cabbage"}
	fmt.Println("Vegetable List:", vegetableList)
	fmt.Println("Length of Vegetable List:", len(vegetableList))

}