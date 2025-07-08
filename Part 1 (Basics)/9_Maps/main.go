package main

import "fmt"


func main()  {
	fmt.Println("Maps in golang")

	// Creating a map using make
	var proggLanguages = make(map[string]string)
	proggLanguages["JS"] = "JavaScript"
	proggLanguages["PY"] = "Python"
	proggLanguages["RB"] = "Ruby"
	proggLanguages["GO"] = "Golang"
	proggLanguages["TS"] = "TypeScript"

	fmt.Println("Programming Languages Map:", proggLanguages)
	fmt.Println("Value for key 'GO':", proggLanguages["GO"])  	// Accessing a value from the map
	fmt.Println("Length of the map:", len(proggLanguages)) 	// Length of the map[type]type

	// Deleting a key-value pair from the map
	delete(proggLanguages, "RB")
	fmt.Println("After deleting 'RB':", proggLanguages)

	// Iterating over the map
	fmt.Println("Iterating over the map:")
	for key, value := range proggLanguages {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}