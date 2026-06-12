package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	fmt.Println("Web requests in golang!")

	res, err := http.Get(url)
	checkNilError(err)

	fmt.Println("Response is of type: ", res)

	defer res.Body.Close() //Always close the response body after using it to avoid memory leaks

	data, err := io.ReadAll(res.Body)
	checkNilError(err)
	fmt.Println("Data inside the response is: ", string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
