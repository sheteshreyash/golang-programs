package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to Shreyash Pizza Palace")
	fmt.Println("Give us your rating out of 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Your rating was", input)
	fmt.Println("Thanks for rating us !")

    // Adding one to the rating
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error converting input to float:", err)
	} else {
		numrating += 1
		fmt.Println("Your rating after adding one is", numrating)
	}
}