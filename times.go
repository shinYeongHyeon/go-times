package times

import (
	"math"
	"time"
)

// GetDateOfFirstMondayOfMonth : get date of first monday
func GetDateOfFirstMondayOfMonth(t time.Time) int {
	dayFirstDateOfMonth := getDayOfFirstDateOfMonth(t)

	if dayFirstDateOfMonth < 2 {
		return 2 - dayFirstDateOfMonth
	}

	return 9 - dayFirstDateOfMonth
}

// GetDateOfFirstTuesdayOfMonth : get date of first tuesday
func GetDateOfFirstTuesdayOfMonth(t time.Time) int {
	dayFirstDateOfMonth := getDayOfFirstDateOfMonth(t)

	if dayFirstDateOfMonth < 3 {
		return 3 - dayFirstDateOfMonth
	}

	return 10 - dayFirstDateOfMonth
}

// GetDateOfFirstWednesdayOfMonth : get date of first wednesday
func GetDateOfFirstWednesdayOfMonth(t time.Time) int {
	dayFirstDateOfMonth := getDayOfFirstDateOfMonth(t)

	if dayFirstDateOfMonth < 4 {
		return 4 - dayFirstDateOfMonth
	}

	return 11 - dayFirstDateOfMonth
}

// GetDateOfFirstThursdayOfMonth : get date of first thursday
func GetDateOfFirstThursdayOfMonth(t time.Time) int {
	dayFirstDateOfMonth := getDayOfFirstDateOfMonth(t)

	if dayFirstDateOfMonth < 5 {
		return 5 - dayFirstDateOfMonth
	}

	return 12 - dayFirstDateOfMonth
}

// GetDateOfFirstFridayOfMonth : get date of first friday
func GetDateOfFirstFridayOfMonth(t time.Time) int {
	dayFirstDateOfMonth := getDayOfFirstDateOfMonth(t)

	if dayFirstDateOfMonth < 6 {
		return 6 - dayFirstDateOfMonth
	}

	return 13 - dayFirstDateOfMonth
}

// GetDateOfFirstSaturdayOfMonth : get date of first saturday
func GetDateOfFirstSaturdayOfMonth(t time.Time) int {
	return 7 - getDayOfFirstDateOfMonth(t)
}

// GetDateOfFirstSaturdayOfMonth : get date of first sunday
func GetDateOfFirstSundayOfMonth(t time.Time) int {
	dayFirstDateOfMonth := getDayOfFirstDateOfMonth(t)

	if dayFirstDateOfMonth == 0 {
		return 1
	}

	return 8 - dayFirstDateOfMonth
}

// GetNthWeekOfMonth : get {n}th Week of Month
// Regard the start of the week as Sunday
func GetNthWeekOfMonth(t time.Time) int {
	date := t.Day()
	daysOfFirstSaturdayOfMonth := GetDateOfFirstSaturdayOfMonth(t)

	if daysOfFirstSaturdayOfMonth >= date {
		return 1
	}

	return 1 + int(math.Ceil(float64(date-daysOfFirstSaturdayOfMonth)/7))
}

// GetLastTimeOfMonth : get Last TimeObject Of Request Month
func GetLastTimeOfMonth(t time.Time) time.Time {
	return getFirstTimeOfMonth(t).AddDate(0, 1, -1)
}

// getDayOfFirstDateOfMonth : get day of first date
func getDayOfFirstDateOfMonth(t time.Time) int {
	return DaysMap[(getFirstTimeOfMonth(t)).Weekday()]
}

// getFirstTimeOfMonth : First time of Month
func getFirstTimeOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
}
