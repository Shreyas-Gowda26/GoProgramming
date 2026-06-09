package main

import (
	"fmt"
	"strconv"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {
		FirstName, LastName, Email, userTickets := UserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUsers(FirstName, LastName, Email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(remainingTickets, conferenceName, bookings []string, FirstName, LastName, Email, userTickets)

			var myMap = make(map[string]string)
			myMap["firstName"] = FirstName
			myMap["lastName"] = LastName
			myMap["email"] = Email
			myMap["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

			getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", getFirstNames())

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidEmail {
				fmt.Println("Your email address is not valid.")
			}
			if !isValidName {
				fmt.Println("Your name is not valid.")
			}
			if !isValidTicketNumber {
				fmt.Println("Your ticket number is not valid.")
			}
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %s booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %d are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func UserInput() (string, string, string, uint) {
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

	return FirstName, LastName, Email, userTickets
}

func bookTickets(remainingTickets uint, conferenceName string, bookings []string, FirstName string, LastName string, Email string, userTickets uint) {
	remainingTickets = uint(remainingTickets) - userTickets
	bookings = append(bookings, myMap)

	fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", FirstName, LastName, userTickets, Email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
