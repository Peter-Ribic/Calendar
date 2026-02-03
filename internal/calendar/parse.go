package calendar

import (
	"strconv"
	"strings"
)

// ParseYear validates and parses a year within bounds.
func ParseYear(s string, minYear int, maxYear int) (int, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, false
	}

	y, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}

	if y < minYear || y > maxYear {
		return 0, false
	}

	return y, true
}
