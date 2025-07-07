package main

import "fmt"

// constants
const LoginToken string = "asdasdasdasdasdasdasd"

func main()  {
	var username string = "Shreyash"
	fmt.Println(username);
	fmt.Printf("Variables is of type %T!\n", username);

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn);
	fmt.Printf("isLoggedIn is of type %T!\n", isLoggedIn);

	var smallVal uint8 = 255
	fmt.Println(smallVal);
	fmt.Printf("smallVal is of type %T!\n", smallVal);

	var smallFloat float32 = 255.5
	fmt.Println(smallFloat);
	fmt.Printf("smallFloat is of type %T!\n", smallFloat);

	var smallFloat2 float64 = 255.5
	fmt.Println(smallFloat2);
	fmt.Printf("smallFloat2 is of type %T!\n", smallFloat2);

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("anotherVariable is of type %T!\n", anotherVariable)

	// implicit type
	var website = "shreyash.com"
	fmt.Println(website)
	fmt.Printf("website is of type %T!\n", website)

	// no var style
	numberOfUsers := 3000  /// walrus operator allowed inside function but not outside eg. global or package level
	fmt.Println(numberOfUsers)
	fmt.Printf("numberOfUsers is of type %T!\n", numberOfUsers)

	// constants
	fmt.Println(LoginToken)
	fmt.Printf("LoginToken is of type %T!\n", LoginToken)
}