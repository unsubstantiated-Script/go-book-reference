package goBook

import (
	"fmt"
	"log"
)

func ToCelsius() {
	fmt.Println("Enter a temp in Fahrenheit:")
	fTemp, err := GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	cTemp := (fTemp - 32) * 5 / 9
	fmt.Printf("%2f degrees Celsius\n", cTemp)
}
