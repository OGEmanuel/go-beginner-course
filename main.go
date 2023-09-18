package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
	// "strconv"
	// "strings"
)

	var conferenceName = "Go Conference"
	// conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	// var bookings [50]string
	// var bookings []string
	// var bookings = make([]map[string]string, 0)
	var bookings = make([]UserData, 0)
	// bookings := []string{}

	type UserData struct {
		firstName string
		lastName  string
		email string
		numberOfTickets int
	}

	var wg = sync.WaitGroup{}

func main() {


	greetUsers()

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Printf("Welcome to %v booking application \n", conferenceName)
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	// fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	// fmt.Println("Get your tickets here to attend")

	// for {
		
firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, int(remainingTickets))

		// isValidCity := city == "Singapore" || city == "London"
		// isInvalidCity := city != "Singapore" || city != "London"

		// if userTickets > int(remainingTickets) {
		// 	fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTickets)
		// 	continue
		// }
	
		if isValidName && isValidEmail && isValidTicketNumber {
			
	bookings = bookTicket(userTickets, firstName, lastName, email)

	wg.Add(1)
	go sendTicket(userTickets, firstName, lastName, email)

		// call funtion print first names
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			// fmt.Println("Your input data is Invalid, try again")
			// continue
		}

		wg.Wait()

		// city := "London"

		// switch city {
		// case "New York":
		// 	// execute code for booking New York conference tickets
		// case "Singapore":
		// 	// execute code for booking Singapore conference tickets
		// case "London", "Berlin":
		// 	// execute code for booking London & Berlin conference tickets
		// // case "Berlin":
		// 	// execute code for booking Berlin conference tickets
		// case "Mexico City":
		// 	// execute code for booking Mexico City conference tickets
		// case "Hong Kong":
		// 	// execute code for booking Hong Kong conference tickets
		// default:
		// 	fmt.Print("No valid city selected")
		// }

		// var noTicketsRemaining bool = remainingTickets == 0
		// noTicketsRemaining := remainingTickets == 0
	// }	
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
		for _, booking := range bookings {
			// var names = strings.Fields(booking)
			// booking["firstname"]
			// var firstName = names[0]
			// firstNames = append(firstNames, booking["firstName"])
			firstNames = append(firstNames, booking.firstName)
		}

		return firstNames
		// fmt.Printf("The first names of bookings are: %v\n", firstNames)
}

func getUserInput() (string, string, string, int) {
	var firstName string
		var lastName string
		var email string
		var userTickets int
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
	
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
	
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) []UserData {
	remainingTickets = remainingTickets - uint(userTickets)

	// create a map for a user
	// var userData = make(map[string]string)
	var UserData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatInt(int64(userTickets), 10)

		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, UserData)
		fmt.Printf("List of bookings is %v\n", bookings)
	
		fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email);
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		return bookings
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}