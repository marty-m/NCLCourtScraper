package main

import "time"

type Booking struct {
	CourtNr int
	Date    time.Time
	Alerted bool
}

func NewBooking(courtNr int, date time.Time) *Booking {
	booking := Booking{CourtNr: courtNr, Date: date}
	booking.Alerted = false
	return &booking
}
