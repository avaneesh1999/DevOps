package main

import (
	"fmt" 
	"strings"
)

func main(){

	//va and const and printing 

	var confrenceName= "PVR"
	const Totaltickets=50
	var RemainingTickets uint
	//slices 
	bookings:=[]string{}

    

	fmt.Printf("Welcome to %v \n",confrenceName);
	fmt.Printf("Get your Tickets here \n");
	
 for{


	var firstName string
	var lastName string
	var email string
	var userTicket uint
	
     //user input & pointers
	fmt.Println("Enter your  first Name ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name:  ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your  email address :")
	fmt.Scan(&email)

	fmt.Println("Enter user Ticket: ")
	fmt.Scan(&userTicket)

	RemainingTickets=Totaltickets-userTicket
	bookings=append(bookings,firstName+" "+lastName)

	
     fmt.Printf(" Thankyou %v %v for Booking %v tickets  you will recieve a confirmation on your %v email address. \n",firstName,lastName,userTicket,email,);
    
	fmt.Printf("There are Total %v Tickets and %v are available now hurry up!!! \n",Totaltickets,RemainingTickets,)
	
	fmt.Printf("These are all your bookiungs %v \n",bookings,)

	// firstNames:=[]string{}
	//   for_, booking := range bookings {
	// 	var names = strings.Fields(booking)
	// 	 firstNames = append(firstNames,names[0])
	//   }
}

    //If else statement
   if RemainingTickets==0{
	  fmt.Printf("OUR TICKETS HAS BEEN BOOKED")
   }

}
