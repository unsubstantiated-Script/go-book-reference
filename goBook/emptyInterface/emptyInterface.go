package emptyInterface

import "fmt"

func acceptAllTheThings(thing interface{}) {
	fmt.Println(thing)
}

func EmptyInterface() {
	acceptAllTheThings(3.1414)
	acceptAllTheThings("Derp")
	acceptAllTheThings(true)
}
