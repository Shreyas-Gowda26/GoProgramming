package main

import "fmt"

func main() {
	//infinite loop
	for {
		fmt.Println("This is an infinite loop")
	}

	//This is a loop
	fmt.Println("E")
	for i := 0; i < 5; i++ {
		fmt.Println()
	}
}
