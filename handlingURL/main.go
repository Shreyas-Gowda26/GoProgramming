package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://github.com/Shreyas-Gowda26"

func main() {
	fmt.Println("Handling URL in golang!")
	fmt.Println(myurl)

	res, err := url.Parse(myurl)
	checkNilError(err)
// Extract and display the scheme, host, and path from the URL.
	fmt.Println("Scheme: ", res.Scheme)
	fmt.Println("Host: ", res.Host)
	fmt.Println("Path: ", res.Path)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
