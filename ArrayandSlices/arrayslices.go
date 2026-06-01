package main

import "fmt"

func main() {
	var FirstName string
	var LastName string

	fmt.Println("Enter the FirstName: ")
	fmt.Scanln(&FirstName)
	fmt.Println("Enter the LastName: ")
	fmt.Scanln(&LastName)

	var arr = [3]string{}
	fmt.Println(arr)
	arr[0] = FirstName + " " + LastName
	fmt.Printf("The whole array is: %v\n", arr)

	var slice []string
	slice = append(slice, "Shreyas")
	fmt.Println(slice)
}
