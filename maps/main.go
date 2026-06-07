package main

import (
	"fmt"
)

func main() {
	var languages = make(map[string]string)

	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["go"] = "Go"

	fmt.Println("The map: ", languages)

	delete(languages, "js")

	fmt.Println("The map after deletion: ", languages)

	for key, value := range languages {
		fmt.Printf("For key  %v the value is %v\n", key, value)
	}
}
