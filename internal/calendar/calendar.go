package calendar

import "time"

// DaysIn returns the number of days in a given month and year.
func DaysIn(month int, year int) int {
	// Day 0 of next month is the last day of this month
	t := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.Local)
	return t.Day()
}

// WeekdayMondayIndex converts Go's weekday (Sunday=0)
// to Monday-based index (Monday=0 ... Sunday=6).
func WeekdayMondayIndex(wd time.Weekday) int {
	return (int(wd) + 6) % 7
}
