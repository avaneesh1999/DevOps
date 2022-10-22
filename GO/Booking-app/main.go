package main


import "fmt"

func main(){
	var confrenceName= "PVR"
	const Totaltickets=50
	var RemainingTickets uint

	fmt.Printf("Welcome to %v \n",confrenceName);
	fmt.Printf("Get your Tickets here \n");
	

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your  first Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last  Name:  ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address :")
	fmt.Scan(&email)

	fmt.Println("Enter user Ticket: ")
	fmt.Scan(&userTicket)

	RemainingTickets=Totaltickets-userTicket

	fmt.Printf(" Thankyou %v %v for Booking %v tickets  you will recieve a confirmation on your %v email address. \n",firstName,lastName,userTicket,email,);
    
	fmt.Printf("there are Total %v Tickets and %v are available now hurry up!!! \n",Totaltickets,RemainingTickets,)





}
