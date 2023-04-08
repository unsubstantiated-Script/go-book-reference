package goBook

import (
	"fmt"
	"goBook/goBook/util"
	"log"
)

func AverageBeef() {
	numbers, err := util.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	var total = sum / float64(len(numbers))

	fmt.Println(total)
}
