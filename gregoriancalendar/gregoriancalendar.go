package gregoriancalendar

import "fmt"

func GregorianCalendar(Day int, Month int, Year int) {
	var totalDays int
	daysInBetweenYears := daysInBetweenAnchorAndTarget(1970, Year)
	daysInTargetYear := daysInTargetYear(Day, Month, Year)
	fmt.Printf("Days in target year: %d\n", daysInTargetYear)
	totalDays = (daysInBetweenYears + daysInTargetYear)
	fmt.Printf("Total Days: %d\n", totalDays)
	result := dayOfWeekToString(calculateDayOfTheWeek(totalDays))
	fmt.Println(result)
}

// Helper functions

/*
Return the number
*/
func calculateDayOfTheWeek(Day int) int {
	anchorDay := 4
	return (Day + anchorDay) % 7

}

/*
Return the name of the weekday
*/
func dayOfWeekToString(dayNumber int) string {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return days[dayNumber]
}

/*
Return the number of days in the target year
*/
func daysInTargetYear(Day int, Month int, Year int) int {
	days := 0
	for i := 0; i <= Month; i++ {
		if i == Month {
			fmt.Printf("In target month: %d\n", i)
			days += Day
		} else {
			if i == 1 || i == 3 || i == 5 || i == 7 || i == 8 || i == 10 || i == 12 {
				days += 31
			} else if i == 4 || i == 6 || i == 9 || i == 11 {
				days += 30
			} else {
				leap := isLeapYear(Year)
				if leap {
					days += 28
				} else {
					days += 29
				}
			}
		}
	}
	return days
}

/*
Return the number of days inbetween the anchor year and target year (int)
*/
func daysInBetweenAnchorAndTarget(AnchorYear int, Year int) int {
	var years int
	var days int
	if Year > AnchorYear {
		years = Year - AnchorYear
	}
	fmt.Printf("Years inbetween 1970 and 2024: %d\n", years)
	days += years * 365
	fmt.Printf("Raw days inbetween 1970 and 2024: %d\n", days)
	days += accountForLeapYears(AnchorYear, Year)
	days = (days - 1)
	fmt.Printf("Days including leap days inbetween 1970 and 2024: %d\n", days)
	return days
}

/*
Add the number of leap days given the range of the years
*/
func accountForLeapYears(AnchorYear int, Year int) int {
	leapDays := 0
	for i := AnchorYear; i <= Year; i++ {
		if isLeapYear(i) {
			leapDays++
		}
	}
	return leapDays
}

/*
Decide if a year is a leap year
*/
func isLeapYear(Year int) bool {
	if (Year%4 == 0 && Year%100 != 100) || Year%400 == 0 {
		return true
	}
	return false
}
