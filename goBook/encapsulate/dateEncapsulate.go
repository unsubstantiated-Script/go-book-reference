package encapsulate

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

//Go Setters
func (d *Date) SetYear(year int) {
	d.Year = year
}
func (d *Date) SetMonth(month int) {
	d.Month = month
}
func (d *Date) SetDay(day int) {
	d.Day = day
}

func Dateige() {
	date := Date{}
	date.SetYear(2022)
	date.SetMonth(4)
	date.SetDay(25)
	fmt.Println(date)
}
