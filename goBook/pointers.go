package goBook

import "fmt"

func Pointers() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)

	myAmountPointer := &amount
	fmt.Println(myAmountPointer)
	fmt.Println(*myAmountPointer)

	*myAmountPointer = 22
	fmt.Println(*myAmountPointer)

	var myFloatyPointer *float64 = createPointer()
	fmt.Println(*myFloatyPointer)
}

//Returning the value of a pointer
func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}
