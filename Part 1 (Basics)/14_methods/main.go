package main

import "fmt"

// Defining a struct
type Person struct {
	firstName string
	lastName  string
	age       int
	Status    bool
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

	shreyash.GetStatus()
	shreyash.newfirstName()
}

// Creating a method for the person struct
func (p Person) GetStatus() {
	fmt.Println("Status of", p.firstName, "is", p.Status)
}

// Creating a method to update the status of the person
func (p *Person) newfirstName() {
	p.firstName = "Omkar"
	fmt.Println("Name of", p.firstName, "is updated to", p.firstName)

}