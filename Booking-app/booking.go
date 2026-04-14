package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("****Welcome to the Go conference booking application****\n")

	const TotalTickets int = 50
	var remainingTickets int = 50
	var firstName string
	var lastName string
	var userEmail string
	var userTickets int
	for {
		// User input reading
		fmt.Printf("Enter you first Name: ")
		fmt.Scan(&firstName)

		fmt.Printf("Enter you Last Name: ")
		fmt.Scan(&lastName)

		fmt.Printf("Enter you Email: ")
		fmt.Scan(&userEmail)

		fmt.Printf("Enter number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		if len(firstName) < 3 || len(lastName) < 3 {
			fmt.Printf("First name and last name should be at least 3 characters long \n")
			continue

		} else if !strings.Contains(userEmail, "@") || !strings.Contains(userEmail, ".") {
			fmt.Printf("Email should contain @ sign and domain \n")
			continue

		} else if userTickets <= 0 || userTickets > remainingTickets {
			fmt.Printf("You can only book between 1 and %v tickets \n", remainingTickets)
			continue
		} else if userTickets < 0 || userTickets > TotalTickets {
			fmt.Printf("You can only book between 1 and %v tickets \n", TotalTickets)
			continue
		}
		fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s \n", firstName, lastName, userTickets, userEmail)

		// Matching remaining tickets with user input
		remainingTickets = remainingTickets - userTickets
		fmt.Printf("%d tickets remaining for the conference \n", remainingTickets)
		if remainingTickets == 0 {
			fmt.Printf("Our conference is booked out. Come back next year.\n")
			break
		}
	}
}
