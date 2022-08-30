package main

import "fmt"

func main() {
	var conferenceName = "DevOps 2022 Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	bookings := []string{"Naman", "ABC"}
	var i = 0

	fmt.Printf("Welcome to Our %v Booking Application \n", conferenceName)
	fmt.Printf("We have total of \n%v Tickets \n& %v Tickets \nare still available\n", remainingTickets, conferenceTickets)
	fmt.Printf("Book Your Tickets Now !!!\n")
	fmt.Println("Enter your First Name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last Name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address")
	fmt.Scan(&email)
	fmt.Println("Enter No. of tickets to be Booked")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	if userTickets >= 2 {
		fmt.Printf("Congratulation %v %v You have successfully Booked %v Tickets. \nYou will receive a confirmation on %v\n", firstName, lastName, userTickets, email)
	}
	fmt.Printf("Total %v bookings are done till now\n By %v", i+1, bookings)
	i++
	fmt.Printf("\nRemaing Tickers for %v are %v", conferenceName, remainingTickets)

}
