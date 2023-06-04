package encapsulate

import (
	"fmt"
	"goBook/goBook/encapsulate/calendar"
	"log"
)

func Dateige() {
	event := calendar.Event{}
	err := event.SetTitle("Momther's Day")
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(4)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(25)
	if err != nil {
		log.Fatal(err)
	}

	//Getting the getters
	fmt.Println(event.Title())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
	fmt.Println(event.Year())

}
