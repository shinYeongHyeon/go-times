package times

import (
	"math"
	"time"
)

// GetDaysOfFirstSaturdayOfMonth : get day of first saturday
// day : DaysMap
func GetDaysOfFirstSaturdayOfMonth(time time.Time) int {
	return 7 - DaysMap[time.Weekday()]
}

// GetNthWeekOfMonth : get {n}th Week of Month
func GetNthWeekOfMonth(time time.Time) int {
	date := time.Day()
	daysOfFirstSaturdayOfMonth := GetDaysOfFirstSaturdayOfMonth(time)

	if daysOfFirstSaturdayOfMonth >= date {
		return 1
	}

	return 1 + int(math.Ceil(float64(date - daysOfFirstSaturdayOfMonth) / 7))
}