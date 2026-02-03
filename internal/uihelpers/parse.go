package uihelpers

import (
	"strconv"
	"strings"
	"time"
)

// ParseDDMMYYYY parses DD.MM.YYYY and validates the date.
func ParseDDMMYYYY(s string, minYear int, maxYear int) (day, month, year int, ok bool) {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return 0, 0, 0, false
	}

	d, e1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	m, e2 := strconv.Atoi(strings.TrimSpace(parts[1]))
	y, e3 := strconv.Atoi(strings.TrimSpace(parts[2]))
	if e1 != nil || e2 != nil || e3 != nil {
		return 0, 0, 0, false
	}

	if y < minYear || y > maxYear || m < 1 || m > 12 || d < 1 || d > 31 {
		return 0, 0, 0, false
	}

	// Validate real date (reject 31.02.x etc.)
	t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	if t.Year() != y || int(t.Month()) != m || t.Day() != d {
		return 0, 0, 0, false
	}

	return d, m, y, true
}
