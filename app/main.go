package main

import (
	"fmt"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %s booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %d are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var FirstName string
	var LastName string
	var Email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scanf("%s", &FirstName)

	fmt.Println("Enter your last name: ")
	fmt.Scanf("%s", &LastName)

	fmt.Println("Enter your email: ")
	fmt.Scanf("%s", &Email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanf("%d", &userTickets)

	remainingTickets = uint(conferenceTickets) - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", FirstName, LastName, userTickets, Email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
