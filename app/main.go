package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
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

		isValidName := len(FirstName) >= 2 && len(LastName) > 2
		isValidEmail := strings.Contains(Email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = uint(remainingTickets) - userTickets

			var bookings []string

			bookings = append(bookings, FirstName+" "+LastName)
			bookings = append(bookings, "Shreyas Chinav")
			fmt.Printf("The whole array: \n%v", bookings)

			fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", FirstName, LastName, userTickets, Email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names are: %v\n", firstNames)

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

func greetUsers(confName string, confTickets int, remTickets uint) {
	fmt.Printf("Welcome to %s booking application!\n", confName)
	fmt.Printf("We have a total of %v tickets and %d are still available.\n", confTickets, remTickets)
	fmt.Println("Get your tickets here to attend")
}
