package main

import (
	"fmt"
	"strings"
)

func main() {
	// var confaranceName = "KubeCon Booking"
	// const totalSeats = 50
	// var remaningTickets = 50
	// fmt.Println("Welcome to", confaranceName, "Booking App")
	// fmt.Printf("We have total %v seat, remaing seats %v \n", totalSeats, remaningTickets)
	//
	// // Data Type is go
	// var userName1 string
	// userName1 = "Sonu"
	// fmt.Printf("Hello %v...!\n", userName1)
	//
	// // Data Type inplace
	// var userName2 string = "Ankita"
	// fmt.Printf("Username is %v \n", userName2)
	//
	// // Data Type
	// userName3 := "Sonu"
	// fmt.Printf("Hello User %v \n", userName3)

	// Booking Ticket Logic
	confranceName := "KubeCon Booking"
	const totalSeats = 50
	var remaningTickets int8 = 50
	fmt.Println("Welcome to", confranceName, "Booking App")
	fmt.Printf("We have total %v seat, remaing seats %v \n", totalSeats, remaningTickets)
	bookings := []string{}
	for {
		var firstName string
		var lastName string
		var userEmail string
		var userTickets uint

		// Array in go only support one data type and fixed size(string only)
		// var bookings = [50]string{"sonu", "bhagat"} -- Manually adding data in array
		// #var bookings [50]string

		// slice in go underlying array and dynamic size and support multiple data type
		// var bookings []string

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&userEmail)

		fmt.Println("Enter How may Ticket you want to book: ")
		fmt.Scan(&userTickets)

		if int8(userTickets) > int8(remaningTickets) {
			fmt.Printf("Sorry, we have only %v tickets remaining. \n", remaningTickets)
			continue
		}

		remaningTickets = int8(remaningTickets) - int8(userTickets)
		//bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("%v %v Thanks for booking %v tickets, You will recived Confirmation email at %v \n", firstName, lastName, userTickets, userEmail)
		fmt.Printf("We have only %v Remaing tickets \n", remaningTickets)

		// For loop inside for loop to get only frst name from slice
		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("First Names of bookings are %v \n", firstNames)

		//fmt.Printf("Printing Slice in list %v\n", bookings)

		if remaningTickets == 0 {
			fmt.Println("Our confrance is booked out. Come back next year.")
			break

		}

	}
	// fmt.Printf("First Slice %v \n", bookings[0])

	//fmt.Printf("Data Type :- %T \n", bookings)
	// fmt.Printf("Lenth of slice %v \n", len(bookings))
}
