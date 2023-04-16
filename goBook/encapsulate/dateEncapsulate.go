package encapsulate

import (
	"fmt"
	"log"
)

func Dateige() {
	date := Date{}
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

	fmt.Println(date)
}
