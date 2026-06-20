package main

import (
	"fmt"
)

func main() {
	fmt.Println("Struct in golang!")

	shreyas := UserData{"Shreyas", "shreyas@example.com", true, 25} //parameters are passed as in the structure
	fmt.Printf("The user data is: %+v\n", shreyas)
}

type UserData struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
