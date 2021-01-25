package times

import (
	"math"
	"time"
)

// GetDaysOfFirstSaturdayOfMonth : will deprecated, recommend GetDateOfFirstSaturdayOfMonth
// day : DaysMap
func GetDaysOfFirstSaturdayOfMonth(t time.Time) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)

	return 7 - DaysMap[firstOfMonth.Weekday()]
}

// GetDateOfFirstMondayOfMonth : get date of first monday
// day : DaysMap
func GetDateOfFirstMondayOfMonth(t time.Time) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)

	firstWeekOfMonth := DaysMap[firstOfMonth.Weekday()]

	if firstWeekOfMonth < 2 {
		return 2 - DaysMap[firstOfMonth.Weekday()]
	}

	return 9 - DaysMap[firstOfMonth.Weekday()]
}

// GetDateOfFirstTuesdayOfMonth : get date of first tuesday
// day : DaysMap
func GetDateOfFirstTuesdayOfMonth(t time.Time) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)

	firstWeekOfMonth := DaysMap[firstOfMonth.Weekday()]

	if firstWeekOfMonth < 3 {
		return 3 - DaysMap[firstOfMonth.Weekday()]
	}

	return 10 - DaysMap[firstOfMonth.Weekday()]
}

// GetDateOfFirstWednesdayOfMonth : get date of first wednesday
// day : DaysMap
func GetDateOfFirstWednesdayOfMonth(t time.Time) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)

	firstWeekOfMonth := DaysMap[firstOfMonth.Weekday()]

	if firstWeekOfMonth < 4 {
		return 4 - DaysMap[firstOfMonth.Weekday()]
	}

	return 11 - DaysMap[firstOfMonth.Weekday()]
}

// GetDateOfFirstThursdayOfMonth : get date of first thursday
// day : DaysMap
func GetDateOfFirstThursdayOfMonth(t time.Time) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)

	firstWeekOfMonth := DaysMap[firstOfMonth.Weekday()]

	if firstWeekOfMonth < 5 {
		return 5 - DaysMap[firstOfMonth.Weekday()]
	}

	return 12 - DaysMap[firstOfMonth.Weekday()]
}

// GetDateOfFirstFridayOfMonth : get date of first friday
// day : DaysMap
func GetDateOfFirstFridayOfMonth(t time.Time) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)

	firstWeekOfMonth := DaysMap[firstOfMonth.Weekday()]

	if firstWeekOfMonth < 6 {
		return 6 - DaysMap[firstOfMonth.Weekday()]
	}

	return 13 - DaysMap[firstOfMonth.Weekday()]
}

// GetDateOfFirstSaturdayOfMonth : get date of first saturday
// day : DaysMap
func GetDateOfFirstSaturdayOfMonth(t time.Time) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)

	return 7 - DaysMap[firstOfMonth.Weekday()]
}

// GetDateOfFirstSaturdayOfMonth : get date of first sunday
// day : DaysMap
func GetDateOfFirstSundayOfMonth(t time.Time) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)

	firstWeekOfMonth := DaysMap[firstOfMonth.Weekday()]

	if firstWeekOfMonth == 0 {
		return 1
	}

	return 8 - DaysMap[firstOfMonth.Weekday()]
}

// GetNthWeekOfMonth : get {n}th Week of Month
// Regard the start of the week as Sunday
func GetNthWeekOfMonth(t time.Time) int {
	date := t.Day()
	daysOfFirstSaturdayOfMonth := GetDaysOfFirstSaturdayOfMonth(t)

	if daysOfFirstSaturdayOfMonth >= date {
		return 1
	}

	return 1 + int(math.Ceil(float64(date-daysOfFirstSaturdayOfMonth)/7))
}
