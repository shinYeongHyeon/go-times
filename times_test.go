package times

import (
	"testing"
	"time"
)

var startedMondayMonth time.Time = time.Date(2021, 2, 3, 0, 0, 0, 0, time.UTC)
var startedSaturdayMonth time.Time = time.Date(2021, 5, 3, 0, 0, 0, 0, time.UTC)
var startedSundayMonth time.Time = time.Date(2021, 8, 3, 0, 0, 0, 0, time.UTC)
var lastTimeOf202108 time.Time = time.Date(2021, 8, 31, 0, 0, 0, 0, time.UTC)
var lastTimeOf202002 time.Time = time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC)
var timeOf20200206 time.Time = time.Date(2020, 2, 6, 0, 0, 0, 0, time.UTC)

func TestGetDateOfFirstMondayOfMonth(t *testing.T) {
	want := 1
	if got := GetDateOfFirstMondayOfMonth(startedMondayMonth); got != want {
		t.Errorf("GetDateOfFirstMondayOfMonth() = %q, want %q", got, want)
	}

	want = 3
	if got := GetDateOfFirstMondayOfMonth(startedSaturdayMonth); got != want {
		t.Errorf("GetDateOfFirstMondayOfMonth() = %q, want %q", got, want)
	}
}

func TestGetDateOfFirstTuesdayOfMonth(t *testing.T) {
	want := 2
	if got := GetDateOfFirstTuesdayOfMonth(startedMondayMonth); got != want {
		t.Errorf("GetDateOfFirstTuesdayOfMonth() = %q, want %q", got, want)
	}

	want = 4
	if got := GetDateOfFirstTuesdayOfMonth(startedSaturdayMonth); got != want {
		t.Errorf("GetDateOfFirstTuesdayOfMonth() = %q, want %q", got, want)
	}
}

func TestGetDateOfFirstWednesdayOfMonth(t *testing.T) {
	want := 3
	if got := GetDateOfFirstWednesdayOfMonth(startedMondayMonth); got != want {
		t.Errorf("GetDateOfFirstWednesdayOfMonth() = %q, want %q", got, want)
	}

	want = 5
	if got := GetDateOfFirstWednesdayOfMonth(startedSaturdayMonth); got != want {
		t.Errorf("GetDateOfFirstWednesdayOfMonth() = %q, want %q", got, want)
	}
}

func TestGetDateOfFirstThursdayOfMonth(t *testing.T) {
	want := 4
	if got := GetDateOfFirstThursdayOfMonth(startedMondayMonth); got != want {
		t.Errorf("GetDateOfFirstThursdayOfMonth() = %q, want %q", got, want)
	}

	want = 6
	if got := GetDateOfFirstThursdayOfMonth(startedSaturdayMonth); got != want {
		t.Errorf("GetDateOfFirstThursdayOfMonth() = %q, want %q", got, want)
	}
}

func TestGetDateOfFirstFridayOfMonth(t *testing.T) {
	want := 5
	if got := GetDateOfFirstFridayOfMonth(startedMondayMonth); got != want {
		t.Errorf("GetDateOfFirstFridayOfMonth() = %q, want %q", got, want)
	}

	want = 7
	if got := GetDateOfFirstFridayOfMonth(startedSaturdayMonth); got != want {
		t.Errorf("GetDateOfFirstFridayOfMonth() = %q, want %q", got, want)
	}
}

func TestGetDateOfFirstSaturdayOfMonth(t *testing.T) {
	want := 6
	if got := GetDateOfFirstSaturdayOfMonth(startedMondayMonth); got != want {
		t.Errorf("GetDaysOfFirstSaturdayOfMonth() = %q, want %q", got, want)
	}
}

func TestGetDateOfFirstSundayOfMonth(t *testing.T) {
	want := 7
	if got := GetDateOfFirstSundayOfMonth(startedMondayMonth); got != want {
		t.Errorf("GetDateOfFirstSundayOfMonth() = %v, want %v", got, want)
	}

	want = 1
	if got := GetDateOfFirstSundayOfMonth(startedSundayMonth); got != want {
		t.Errorf("GetDateOfFirstSundayOfMonth() = %v, want %v", got, want)
	}
}

func TestGetNthWeekOfMonth(t *testing.T) {
	want := 1
	if got := GetNthWeekOfMonth(startedMondayMonth); got != want {
		t.Errorf("GetNthWeekOfMonth() = %v, want %v", got, want)
	}

	want = 2
	if got := GetNthWeekOfMonth(startedSaturdayMonth); got != want {
		t.Errorf("GetNthWeekOfMonth() = %v, want %v", got, want)
	}
}

func TestGetLastTimeOfMonth(t *testing.T) {
	want := lastTimeOf202108
	if got := GetLastTimeOfMonth(startedSundayMonth); got != want {
		t.Errorf("GetLastTimeOfMonth() = %v, want %v", got, want)
	}

	want = lastTimeOf202002
	if got := GetLastTimeOfMonth(timeOf20200206); got != want {
		t.Errorf("GetLastTimeOfMonth() = %v, want %v", got, want)
	}
}