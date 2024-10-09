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
	var bookings [10]string // array to store the bookings names, size is 10, type is string
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName) // fmt.Scan() is used to take input from the user
	fmt.Print("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTickets)                                  // & is used to get the address of the variable
	remainingTickets = remainingTickets - uint(userTickets) // typecasting userTickets to uint as remainingTickets is of type uint, or change userTickets to uint
	bookings[0] = userName
	bookings[1] = "Test User 2"
	bookings[2] = "Test User 3"
	bookings[3] = "Test User 4"
	fmt.Printf("User %v is booking %v tickets\nTickets remaining : %v\n", userName, userTickets, remainingTickets)
	fmt.Println("Bookings so far: ", bookings)
	fmt.Println("Booking[0] : ", bookings[0])
	fmt.Println("Thank you for booking !!")

}
