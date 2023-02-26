package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"regexp"
	"time"
)

// SSO sign in page: https://appspay.ncl.ac.uk/Shibboleth.sso/SPNEGO?target=https%3a%2f%2fappspay.ncl.ac.uk/sport/booking

func Court5() {
	//Declare variables to hold results of JS expression evaluation blocks
	var jsres int
	var throwaway []byte
	var rawCell string
	//Create browser instance
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
	defer cancel()
	//Sequence to navigate to Court 5 page on booking site
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://appspay.ncl.ac.uk/Shibboleth.sso/SPNEGO?target=https%3a%2f%2fappspay.ncl.ac.uk/sport/booking"),
		chromedp.WaitReady("body"),
		chromedp.SendKeys("input[name=j_username]", "---NCL username---"),
		chromedp.SendKeys("input[name=j_password]", "---NCL password---"),
		chromedp.Click("button[name=_eventId_proceed]", chromedp.NodeVisible),
		chromedp.WaitReady("head[id=ctl00_Head1]"),
		chromedp.Navigate("https://sportsbookings.ncl.ac.uk/Connect/mrmselectsite.aspx?disableSiteSelection=1"),
		chromedp.Click("input[value=\"Sports Hall 1\"", chromedp.NodeVisible),
		chromedp.Click("input[value=\"Basketball Hoop Court 5\"", chromedp.NodeVisible),
		chromedp.WaitReady("head[id=ctl00_Head1]"),
	)

	// For loop cycling 3 pages of the calendar (max time in advance that court times are shown on app)
	for i := 0; i < 3; i++ {
		err = chromedp.Run(ctx,
			chromedp.Sleep(time.Second*1),
			chromedp.Evaluate("document.getElementsByClassName('itemavailable').length", &jsres),
			chromedp.Evaluate("courtObjArr = document.getElementsByClassName('itemavailable')", &throwaway),
			chromedp.Click("button[data-qa-id=\"label-dateForward\"]"),
		)
		//Check if any open slots are available
		if jsres != 0 {
			//Iterate over open slots
			for i := 0; i < jsres; i++ {
				err = chromedp.Run(ctx,
					//Extract attribute from the slot containing the date & time of slot
					chromedp.Evaluate(fmt.Sprintf("document.getElementsByClassName('itemavailable')[%d].querySelector('span > input').getAttribute('data-qa-id')", i), &rawCell),
				)
				//RegEx searching for date and time respectively
				reDate := regexp.MustCompile("([A-Za-z0-9]+(/[A-Za-z0-9]+)+).*[0-9]{2}:[0-9]{2}:[0-9]{2}(\\.[0-9]{1,3})?")

				//Format date and time strings into 'time' type vars
				courtDate, _ := time.Parse("02/01/2006 15:04:05", reDate.FindString(rawCell))
				fmt.Println(*NewBooking(5, courtDate))
				AddBooking(*NewBooking(5, courtDate))

			}

		}

	}

	if err != nil {
		log.Fatal(err)
	}
}
