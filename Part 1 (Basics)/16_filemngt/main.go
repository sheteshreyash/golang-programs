package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)


func main()  {
	fmt.Println("File Management System in golang")
	content := "This is a simple file management system written in Go."

	// Create a file
	file, err := os.Create("./file.txt")
	if err != nil {
		panic(err)
	}

	// Write to the file
	len, err := io.WriteString(file, content)
	checkNilerror(err)
	fmt.Printf("Wrote %d bytes to file.txt\n", len)

	// Close the file
	defer file.Close()

	// Read from the file
	readFile("./file.txt")

	// Delete
	// deleteFile("./file.txt")
}

// Read from the file
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilerror(err)
	fmt.Println("File content:")
	fmt.Println(string(databyte))

}

// Delete the file
// func deleteFile(filename string) {
// 	err := os.Remove(filename)
//     checkNilerror(err)
// 	fmt.Printf("File %s deleted successfully\n", filename)
// }

// Declaring the error function
func checkNilerror(err error) {
	if err != nil {
		panic(err)
	}
}