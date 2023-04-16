package calendar

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

// Go Setters. Do Not export by keeping struct info lowercase
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

// Go Getters...heh
func (d *Date) Year() int {
	return d.year
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Day() int {
	return d.day
}
