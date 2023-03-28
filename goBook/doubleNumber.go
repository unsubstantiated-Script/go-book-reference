package goBook

import "fmt"

func DoubleNumber() {
	amount := 6
	doubleNumber(&amount)
	fmt.Println(amount)
}

func doubleNumber(number *int) {
	*number *= 2
}

func negate(myBool *bool) {
	*myBool = !*myBool
}

func NegateBool() {
	truth := true
	fmt.Println("Truth is", truth)
	negate(&truth)
	fmt.Println("Truth is now", truth)
	lies := false
	fmt.Println("Lies are", lies)
	negate(&lies)
	fmt.Println("Lies are now", lies)
	fmt.Println(lies)
}
