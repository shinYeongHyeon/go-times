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
var MonthMapByInt = map[int]time.Month{
	1: time.January,
	2: time.February,
	3: time.March,
	4: time.April,
	5: time.May,
	6: time.June,
	7: time.July,
	8: time.August,
	9: time.September,
	10: time.October,
	11: time.November,
	12: time.December,
}
