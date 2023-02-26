package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"time"
)

func SendAlert() {
	msgBody := "The following courts are available:\n"
	addText := ""
	delAmount := 0
	for i := 0; i < len(AvailableBookings); i++ {
		book := &AvailableBookings[i]

		if book.Date.Before(time.Now()) {
			AvailableBookings = append(AvailableBookings[:i], AvailableBookings[i+1:]...)
			delAmount++
			i -= delAmount
			continue
		} else if book.Alerted == false {
			fmtDate := book.Date.Format("(15:04) Monday, 02 January")
			addText += fmt.Sprintf("[Court %v]\t %v\n\n", book.CourtNr, fmtDate)
			book.Alerted = true

		}
	}
	msgBody += addText
	//If no valid (in the future and not yet alerted) new bookings are found, exit function and don't alert.
	if addText == "" {
		fmt.Println("No new available slots found, trying again in 2 hours...")
		return
	}
	//Twilio integration
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: "---twilio account sid---",
		Password: "---twilio api token---",
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo("---recipient phone#---")
	params.SetFrom("---twilio phone#---")
	params.SetBody(msgBody)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!\nWill check again in 2 hours.")
	}
}
