package countVotes

import (
	"embed"
	_ "embed"
	"fmt"
	"goBook/goBook/util"
	"log"
)

//go:embed votes.txt
var contents embed.FS

func CountVotes() {
	lines, err := util.GetStrings(contents, "votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}

}
