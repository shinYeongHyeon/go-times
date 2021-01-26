package times

import (
	"testing"
	"time"
)

var exampleDate time.Time = time.Date(2021, 2, 3, 0, 0, 0, 0, time.UTC)

func TestGetDateOfFirstMondayOfMonth(t *testing.T) {
	want := 1
	if got := GetDateOfFirstMondayOfMonth(exampleDate); got != want {
		t.Errorf("GetDateOfFirstMondayOfMonth() = %q, want %q", got, want)
	}
}

func TestGetDateOfFirstTuesdayOfMonth(t *testing.T) {
	want := 2
	if got := GetDateOfFirstTuesdayOfMonth(exampleDate); got != want {
		t.Errorf("GetDateOfFirstTuesdayOfMonth() = %q, want %q", got, want)
	}
}

func TestGetDateOfFirstWednesdayOfMonth(t *testing.T) {
	want := 3
	if got := GetDateOfFirstWednesdayOfMonth(exampleDate); got != want {
		t.Errorf("GetDateOfFirstWednesdayOfMonth() = %q, want %q", got, want)
	}
}

func TestGetDateOfFirstThursdayOfMonth(t *testing.T) {
	want := 4
	if got := GetDateOfFirstThursdayOfMonth(exampleDate); got != want {
		t.Errorf("GetDateOfFirstThursdayOfMonth() = %q, want %q", got, want)
	}
}

func TestGetDateOfFirstFridayOfMonth(t *testing.T) {
	want := 5
	if got := GetDateOfFirstFridayOfMonth(exampleDate); got != want {
		t.Errorf("GetDateOfFirstFridayOfMonth() = %q, want %q", got, want)
	}
}

func TestGetDateOfFirstSaturdayOfMonth(t *testing.T) {
	want := 6
	if got := GetDateOfFirstSaturdayOfMonth(exampleDate); got != want {
		t.Errorf("GetDaysOfFirstSaturdayOfMonth() = %q, want %q", got, want)
	}
}

func TestGetDateOfFirstSundayOfMonth(t *testing.T) {
	want := 7
	if got := GetDateOfFirstSundayOfMonth(exampleDate); got != want {
		t.Errorf("GetDateOfFirstSundayOfMonth() = %v, want %v", got, want)
	}
}

func TestGetNthWeekOfMonth(t *testing.T) {
	want := 1
	if got := GetNthWeekOfMonth(exampleDate); got != want {
		t.Errorf("GetNthWeekOfMonth() = %v, want %v", got, want)
	}
}
