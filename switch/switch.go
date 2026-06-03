package main

import "fmt"

func main() {
	series := "suits"

	switch series {
	case "breaking bad":
		fmt.Println("You are watching Breaking Bad")

	case "suits":
		fmt.Println("You are watching Suits")

	case "game of thrones":
		fmt.Println("You are watching Game of Thrones")

	default:
		fmt.Println("You are watching something else")
	}
}
