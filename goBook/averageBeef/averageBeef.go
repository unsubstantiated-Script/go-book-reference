package averageBeef

import (
	"embed"
	_ "embed"
	"fmt"
	"goBook/goBook/util"
	"log"
)

//go:embed beefPrices.txt
var contents embed.FS

func AverageBeef() {
	numbers, err := util.GetFloats(contents, "beefPrices.txt")
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
