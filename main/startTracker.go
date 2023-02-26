package main

import (
	"fmt"
	"time"
)

var AvailableBookings []Booking

func AddBooking(booking Booking) {
	//Check if booking was already previously found and alerted
	for i := 0; i < len(AvailableBookings); i++ {
		book := AvailableBookings[i]
		if book.Date == booking.Date && book.CourtNr == booking.CourtNr {
			return
		}
	}
	AvailableBookings = append(AvailableBookings, booking)

}

func main() {
	AvailableBookings = []Booking{{0, time.Now(), true}}
	for true {
		fmt.Println(fmt.Sprintf("---------%s---------", time.Now().Format("2006-01-02  15:04:05")))
		fmt.Println("Starting alert service...")
		fmt.Println("Searching available slots for Court 5...")
		Court5()
		fmt.Println("Searching available slots for Court 8...")
		Court8()
		SendAlert()
		time.Sleep(2 * time.Hour)
	}

}
