package encapsulate

import (
	"fmt"
	"goBook/goBook/encapsulate/calendar"
	"log"
)

func Dateige() {
	date := calendar.Date{}
	err := date.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(4)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(25)
	if err != nil {
		log.Fatal(err)
	}

	//Getting the getters
	fmt.Println(date.Month())
	fmt.Println(date.Day())
	fmt.Println(date.Year())
}
