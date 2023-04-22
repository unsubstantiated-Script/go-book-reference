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

// Default print format
func (m MilliLiters) String() string {
	return fmt.Sprintf("%0.2f milliters", m)
}

type Liters float64

// Default print format
func (l Liters) String() string {
	return fmt.Sprintf("%0.2f liters", l)
}

type Gallons float64

// Default print format
func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gallons", g)
}

//Conversions
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

//Main function
func ConvertLiquidVolume() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())

	water := MilliLiters(500)
	fmt.Printf("%0.3f mililiters equals %0.3f gallons\n", soda, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliLiters())

	// Printing out some default string formats
	fmt.Println(Gallons(12.09248342))
	fmt.Println(Liters(12.09248342))
	fmt.Println(MilliLiters(12.09248342))
}

