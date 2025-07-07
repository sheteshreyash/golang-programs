package main

import "fmt"

// Defining a struct
type Person struct {
	firstName string
	lastName  string
	age       int
}

func main()  {
	fmt.Println("structs in golang")

	// Creating an instance of Person struct
	shreyash := Person{firstName:"Shreyash", lastName:"Kumar", age: 25,}
	fmt.Println(shreyash)
	fmt.Printf("First Name: %s, Last Name: %s, Age: %d\n", shreyash.firstName, shreyash.lastName, shreyash.age)
	fmt.Printf("Type of shreyash: %T\n", shreyash)

	// Accessing fields of the struct
	fmt.Println("First Name:", shreyash.firstName)
	fmt.Println("Last Name:", shreyash.lastName)
	fmt.Println("Age:", shreyash.age)

	// Modifying fields of the struct
	shreyash.age = 26
	fmt.Println("Updated Age:", shreyash.age)
	fmt.Println("Updated Person:", shreyash)

	// Creating another instance of Person struct
	anotherPerson := Person{firstName: "John", lastName: "Doe", age: 30}
	fmt.Println("Another Person:", anotherPerson)
	fmt.Printf("Type of anotherPerson: %T\n", anotherPerson)

	// Accessing fields of anotherPerson
	fmt.Println("First Name:", anotherPerson.firstName)
	fmt.Println("Last Name:", anotherPerson.lastName)
	fmt.Println("Age:", anotherPerson.age)

	// Modifying fields of anotherPerson
	anotherPerson.age = 31
	fmt.Println("Updated Age of Another Person:", anotherPerson.age)
	fmt.Println("Updated Another Person:", anotherPerson)

}
