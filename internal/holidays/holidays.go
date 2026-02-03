package holidays

// Holiday represents one holiday entry.
// If Recurring is true, the holiday repeats every year on Day.Month.
// If Recurring is false, it only applies to the exact Year.
type Holiday struct {
	Day       int
	Month     int
	Year      int
	Recurring bool
}

func IsHoliday(day, month, year int, list []Holiday) bool {
	for _, h := range list {
		if h.Day != day || h.Month != month {
			continue
		}
		if h.Recurring || h.Year == year {
			return true
		}
	}
	return false
}
