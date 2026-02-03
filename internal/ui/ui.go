package ui

import (
	"image/color"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	MinYear = 1
	MaxYear = 9999
)

// Run creates the window with month/year controls and refreshes the manual grid on change.
func Run(a fyne.App) {
	w := a.NewWindow("Calendar")

	months := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	now := time.Now()
	selectedMonth := int(now.Month())
	selectedYear := now.Year()

	monthSelect := widget.NewSelect(months, nil)
	monthSelect.SetSelected(months[selectedMonth-1])

	yearEntry := widget.NewEntry()
	yearEntry.SetText(strconv.Itoa(selectedYear))
	yearEntry.SetPlaceHolder("Year")

	status := widget.NewLabel("")

	//holidayList, _ := holidays.Load("holidays.txt", MinYear, MaxYear)

	weekdayHeader := container.NewGridWithColumns(7,
		headerLabel("Mon"),
		headerLabel("Tue"),
		headerLabel("Wed"),
		headerLabel("Thu"),
		headerLabel("Fri"),
		headerLabel("Sat"),
		headerLabel("Sun"),
	)

	cellTexts := make([]*canvas.Text, 42)
	cellBgs := make([]*canvas.Rectangle, 42)
	cells := make([]fyne.CanvasObject, 0, 42)

	for i := 0; i < 42; i++ {
		bg := canvas.NewRectangle(colNormal())
		bg.SetMinSize(fyne.NewSize(44, 34))

		txt := canvas.NewText("", color.Black)
		txt.Alignment = fyne.TextAlignCenter
		txt.TextSize = 14

		cellBgs[i] = bg
		cellTexts[i] = txt
		cells = append(cells, container.NewMax(bg, container.NewCenter(txt)))
	}

	grid := container.NewGridWithColumns(7, cells...)

	render := func(month int, year int) {
		for i := 0; i < 42; i++ {
			cellTexts[i].Text = ""
			cellBgs[i].FillColor = colNormal()
			cellBgs[i].Refresh()
			cellTexts[i].Refresh()
		}

		first := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
		offset := weekdayMondayIndex(first.Weekday())
		days := daysIn(month, year)

		for d := 1; d <= days; d++ {
			idx := offset + (d - 1)
			if idx < 0 || idx >= 42 {
				continue
			}
			cellTexts[idx].Text = strconv.Itoa(d)

			col := idx % 7
			if col == 6 { // Sunday in Mon..Sun layout
				cellBgs[idx].FillColor = colSunday()
				cellBgs[idx].Refresh()
			}

			cellTexts[idx].Refresh()
		}
	}

	monthSelect.OnChanged = func(s string) {
		m := indexOf(months, s) + 1
		if m < 1 || m > 12 {
			return
		}
		selectedMonth = m

		y, ok := parseYear(yearEntry.Text)
		if !ok {
			status.SetText("Invalid year.")
			return
		}
		status.SetText("")
		selectedYear = y
		render(selectedMonth, selectedYear)
	}

	yearEntry.OnChanged = func(s string) {
		y, ok := parseYear(s)
		if !ok {
			status.SetText("Invalid year.")
			return
		}
		status.SetText("")
		selectedYear = y
		render(selectedMonth, selectedYear)
	}

	render(selectedMonth, selectedYear)

	controls := container.NewHBox(
		widget.NewLabel("Month:"),
		monthSelect,
		widget.NewLabel("Year:"),
		yearEntry,
	)

	w.SetContent(container.NewVBox(
		controls,
		weekdayHeader,
		grid,
		status,
	))

	w.Resize(fyne.NewSize(420, 380))
	w.ShowAndRun()
}

func headerLabel(s string) fyne.CanvasObject {
	t := canvas.NewText(s, color.Black)
	t.TextStyle = fyne.TextStyle{Bold: true}
	t.Alignment = fyne.TextAlignCenter
	t.TextSize = 13

	bg := canvas.NewRectangle(color.RGBA{240, 240, 240, 255})
	bg.SetMinSize(fyne.NewSize(44, 28))

	return container.NewMax(bg, container.NewCenter(t))
}

func colNormal() color.Color { return color.RGBA{255, 255, 255, 255} }

func daysIn(month int, year int) int {
	t := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.Local)
	return t.Day()
}

func weekdayMondayIndex(wd time.Weekday) int { return (int(wd) + 6) % 7 }

func parseYear(s string) (int, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, false
	}
	y, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	if y < MinYear || y > MaxYear {
		return 0, false
	}
	return y, true
}

func indexOf(list []string, value string) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}

func colSunday() color.Color { return color.RGBA{255, 235, 235, 255} }
