package holidays

import (
	"bufio"
	"os"
	"strings"
	"unicode"

	"github.com/Peter-Ribic/Calendar/internal/uihelpers"
)

// Load reads holidays from an ASCII file.
// Line format: DD.MM.YYYY <delimiter> recurringFlag
// Delimiter can be any non-alphanumeric char ('.' is kept so the date stays intact).
func Load(path string, minYear int, maxYear int) ([]Holiday, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out []Holiday
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		tokens := splitTokens(line)
		if len(tokens) < 2 {
			continue
		}

		d, m, y, ok := uihelpers.ParseDDMMYYYY(tokens[0], minYear, maxYear)
		if !ok {
			continue
		}
		rec, ok := parseRecurringFlag(tokens[1])
		if !ok {
			continue
		}

		out = append(out, Holiday{Day: d, Month: m, Year: y, Recurring: rec})
	}

	return out, nil
}

func splitTokens(line string) []string {
	return strings.FieldsFunc(line, func(r rune) bool {
		if r == '.' {
			return false
		}
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
}

func parseRecurringFlag(s string) (bool, bool) {
	s = strings.ToLower(strings.TrimSpace(s))
	switch s {
	case "1", "true", "yes", "y", "r":
		return true, true
	case "0", "false", "no", "n", "nr":
		return false, true
	default:
		return false, false
	}
}
