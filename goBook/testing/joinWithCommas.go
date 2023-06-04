package testing

import (
	"fmt"
	"strings"
)

func BuildJoin() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", joinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", joinWithCommas(phrases))
}

func joinWithCommas(phrases []string) string {
	//Appending commas to the string
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	result += ", and "
	result += phrases[len(phrases)-1]
	return result
}
