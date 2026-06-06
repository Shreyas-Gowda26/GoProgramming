package main

import "strings"

func validateUsers(FirstName string, LastName string, Email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(FirstName) >= 2 && len(LastName) > 2
	isValidEmail := strings.Contains(Email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
