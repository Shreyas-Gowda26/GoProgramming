package main

import "fmt"

func main() {
	fmt.Println("Lets complete the tasks!")
	var task1 string = "Learn devops"
	var task2 string = "Complete go programming"
	const task3 int = 3

	fmt.Printf("Try to %s and try to %s within %d days\n", task1, task2, task3)

	task := "writing"

	switch task {
	case "reading":
		fmt.Println("You are reading a book")

	case "swimming":
		fmt.Println("You are in a pool")

	case "running":
		fmt.Println("You are running")

	case "writing":
		fmt.Println("You are writing")

	default:
		fmt.Println("You are doing nothing")
	}

}
