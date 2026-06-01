package main

import "fmt"

func main() {
	//infinite loop
	//for {
	//	fmt.Println("This is an infinite loop")
	//}

	var a1 []string

	a1 = append(a1, "Shreyas")
	a1 = append(a1, "Chinav")
	a1 = append(a1, "Go")

	//This is a loop
	for _, arr := range a1 {
		fmt.Println(arr)
	}
}
