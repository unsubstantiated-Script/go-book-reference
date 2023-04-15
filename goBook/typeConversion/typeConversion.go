package typeConversion

import "fmt"

type Number int

func (n *Number) Display() {
	//Pointer is needed here
	fmt.Println(*n)
}

func (n *Number) Double() {
	//Pointer is needed here
	*n *= 2
}

func DoubleTypeNumber() {
	number := Number(4)

	number.Display()
	number.Double()
	number.Display()
}

/**
* Converting Liguid Types with GoLang Types
 */
type MilliLiters float64
type Liters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m MilliLiters) ToGallons() Gallons {
	return Gallons(m * .000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliLiters() MilliLiters {
	return MilliLiters(g * 3785.41)
}

func ConvertLiquidVolume() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())

	water := MilliLiters(500)
	fmt.Printf("%0.3f mililiters equals %0.3f gallons\n", soda, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliLiters())
}

