package main

import "fmt"

func main() {
	const totalTickets = 50
	var remainingTickets = totalTickets
	var conferenceName = "Test Conference"
	fmt.Println("Welcome to the ", conferenceName, " Booking System")
	fmt.Println("We have ", totalTickets, " slots total with only ", remainingTickets, "tickets remaining !!")

}
