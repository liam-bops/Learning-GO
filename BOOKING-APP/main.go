package main

import "fmt"

func main() {
	const totalTickets = 50                  // constant value for total tickets, cannot be changed, go auto detects the type when we assign a value
	var remainingTickets uint = totalTickets // variable to store the remaining tickets, uint is used to store only positive integers(enforces +ve constraint)
	conferenceName := "Test Conference"      // := is used to declare and assign a value to a variable, go auto detects the type when we assign a value

	fmt.Printf("conferenceName is of type %T\n", conferenceName) // %T is used to print the type of the variable
	fmt.Printf("totalTickets is of type %T\n", totalTickets)
	fmt.Printf("remainingTickets is of type %T\n", remainingTickets)

	fmt.Println("Welcome to the ", conferenceName, " Booking System")                                        // print similar to python using , as separator
	fmt.Printf("We have %v slots total with only %v tickets remaining !!\n", totalTickets, remainingTickets) // %v is used as a placeholder to print the value of the variable
	fmt.Print("Get your tickets now !!\n")

	var userName string
	var userTickets int

	userName = "John Doe"
	userTickets = 5
	fmt.Printf("User %v is booking %v tickets\n", userName, userTickets)

}
