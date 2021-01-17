package times

import (
	"math"
	"time"
)

// GetDaysOfFirstSaturdayOfMonth : get day of first saturday
// day : DaysMap
func GetDaysOfFirstSaturdayOfMonth(t time.Time) int {
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	return 7 - DaysMap[firstOfMonth.Weekday()]
}

// GetNthWeekOfMonth : get {n}th Week of Month
func GetNthWeekOfMonth(t time.Time) int {
	date := t.Day()
	daysOfFirstSaturdayOfMonth := GetDaysOfFirstSaturdayOfMonth(t)

	if daysOfFirstSaturdayOfMonth >= date {
		return 1
	}

	return 1 + int(math.Ceil(float64(date - daysOfFirstSaturdayOfMonth) / 7))
}