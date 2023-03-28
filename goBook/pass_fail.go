package goBook

import (
	"fmt"
	"log"
)

func PassFail() {
	fmt.Println("Enter a grade:")
	grade, err := GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	var status string

	if grade >= 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
