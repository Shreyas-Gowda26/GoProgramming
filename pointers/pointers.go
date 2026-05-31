package main

import "fmt"

func main() {
	fmt.Print("Pointers in go\n")
	var name string = "Shreyas"
	fname := &name
	fmt.Println("Enter the name of yours: ")
	fmt.Scanln(&name)
	fmt.Println("Hello", name)
	fmt.Println("The value of fname is: ", *fname)
}
