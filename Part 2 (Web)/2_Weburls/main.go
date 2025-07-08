package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?course=golang&paymentid=12345"

func main()  {
	fmt.Println("Handling web URLs in Go")
	fmt.Println("URL:", myurl)

	// parsing the URL
	result, _ := url.Parse(myurl)

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("Raw Query:", result.RawQuery)
	fmt.Println("Query Parameters:", result.Query())

	// Getting specific query parameters
	qparams := result.Query()
	fmt.Printf("Type of query parameters: %T\n", qparams)
	fmt.Println("Course:", qparams["course"][0])

	for key, value := range qparams {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}

	// parts of the URL
	partsOfURL := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "/tutcss",
		RawQuery: "course=reactjs&paymentid=555",
		RawPath:  "user=shreyash&age=20",
	}
	fmt.Println("Constructed URL:", partsOfURL.String())
}