package main

import (
	"go-playground/go-playground/gregoriancalendar"
)

func main() {
	var day int
	var month int
	var year int
	day = 8
	month = 10
	year = 2024

	gregoriancalendar.GregorianCalendar(day, month, year)
}
