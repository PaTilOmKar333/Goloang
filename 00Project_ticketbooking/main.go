package main

import "fmt"

func main() {
	const confernaceTickets = 50
	confernaceName := "GO Confernace"
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", confernaceName)
	fmt.Printf("we have total %v tickets and now availble tickets are %v\n", confernaceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your firstname:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your lastname:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter no of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets.you will recieve mail on your %v mail address. Now %v tickets are remaining.\n", firstName, lastName, userTickets, email, remainingTickets)
		fmt.Printf("all bookings are %v\n", bookings)
	}

}
