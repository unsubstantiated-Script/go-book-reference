package numFloats

import (
	"fmt"
	"log"
	"os"
)

func SumIt() {
	//This gonna read the file name from the CLI
	numbers, err := GetFloats(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Sum: %0.2f\n", sum)
}
