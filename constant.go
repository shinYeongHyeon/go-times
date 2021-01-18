package times

import "time"

// DaysMap is dependent for time package
var DaysMap = map[time.Weekday]int{
	time.Sunday:    0,
	time.Monday:    1,
	time.Tuesday:   2,
	time.Wednesday: 3,
	time.Thursday:  4,
	time.Friday:    5,
	time.Saturday:  6,
}

// MonthMap is dependent for time package
var MonthMap = map[time.Month]int{
	time.January:   1,
	time.February:  2,
	time.March:     3,
	time.April:     4,
	time.May:       5,
	time.June:      6,
	time.July:      7,
	time.August:    8,
	time.September: 9,
	time.October:   10,
	time.November:  11,
	time.December:  12,
}
