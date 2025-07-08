package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://go.dev/doc/"

func main()  {
	fmt.Println("Handling web requests in Go")

	// Here we would typically use the net/http package to make web requests
	response, err := http.Get(url)
	checkError(err)
	fmt.Printf("Type of response is %T\n", response)
	defer response.Body.Close() // Ensure the response body is closed after we're done with it

	// Reading the response
	readResponse(response)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// reading the reponse
func readResponse(response *http.Response) {
	databytes, err := ioutil.ReadAll(response.Body)
	checkError(err)
	fmt.Println("Response body:", string(databytes))
	fmt.Println("Response status code:", response.StatusCode)
	fmt.Println("Response status:", response.Status)
	fmt.Println("Response headers:", response.Header)
	fmt.Println("Response content length:", response.ContentLength)
	
}
