package ui

import (
	"image/color"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Peter-Ribic/Calendar/internal/calendar"
	"github.com/Peter-Ribic/Calendar/internal/holidays"
	"github.com/Peter-Ribic/Calendar/internal/uihelpers"
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

	// Define ui elements.

	monthSelect := widget.NewSelect(months, nil)
	monthSelect.SetSelected(months[selectedMonth-1])

	yearEntry := widget.NewEntry()
	yearEntry.SetText(strconv.Itoa(selectedYear))
	yearEntry.SetPlaceHolder("Year")

	jumpEntry := widget.NewEntry()
	jumpEntry.SetPlaceHolder("Jump to date (DD.MM.YYYY) and press Enter")

	status := widget.NewLabel("")

	// Client did not request nor pay for error handling :).
	holidayList, _ := holidays.Load("holidays.txt", MinYear, MaxYear)

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

	// Create grid cells.
	for i := 0; i < 42; i++ {
		cellBgs[i] = canvas.NewRectangle(colNormal())
		cellBgs[i].SetMinSize(fyne.NewSize(44, 34))

		cellTexts[i] = canvas.NewText("", color.Black)
		cellTexts[i].Alignment = fyne.TextAlignCenter
		cellTexts[i].TextSize = 14

		cells = append(cells, container.NewStack(cellBgs[i], container.NewCenter(cellTexts[i])))
	}

	grid := container.NewGridWithColumns(7, cells...)

	render := func(month int, year int) {
		// Reset all cells.
		for i := 0; i < 42; i++ {
			cellTexts[i].Text = ""
			cellBgs[i].FillColor = colNormal()
			cellBgs[i].Refresh()
			cellTexts[i].Refresh()
		}

		first := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
		offset := calendar.WeekdayMondayIndex(first.Weekday())
		days := calendar.DaysIn(month, year)

		for d := 1; d <= days; d++ {
			idx := offset + (d - 1)
			if idx < 0 || idx >= 42 {
				// Safety check, should not happen.
				continue
			}
			cellTexts[idx].Text = strconv.Itoa(d)

			col := idx % 7
			isSunday := (col == 6)
			isHol := holidays.IsHoliday(d, month, year, holidayList)

			switch {
			case isHol && isSunday:
				cellBgs[idx].FillColor = colHolidaySunday()
			case isHol:
				cellBgs[idx].FillColor = colHoliday()
			case isSunday:
				cellBgs[idx].FillColor = colSunday()
			default:
				cellBgs[idx].FillColor = colNormal()
			}
			cellBgs[idx].Refresh()
			cellTexts[idx].Refresh()
		}
	}

	monthSelect.OnChanged = func(s string) {
		selectedMonth := indexOf(months, s) + 1
		if selectedMonth < 1 || selectedMonth > 12 {
			// Should not happen.
			return
		}
		selectedYear, ok := calendar.ParseYear(yearEntry.Text, MinYear, MaxYear)
		if !ok {
			status.SetText("Invalid year.")
			return
		}
		status.SetText("")
		render(selectedMonth, selectedYear)
	}

	yearEntry.OnChanged = func(s string) {
		selectedYear, ok := calendar.ParseYear(s, MinYear, MaxYear)
		if !ok {
			status.SetText("Invalid year.")
			return
		}
		status.SetText("")
		render(selectedMonth, selectedYear)
	}

	jumpEntry.OnSubmitted = func(s string) {
		_, selectedMonth, selectedYear, ok := uihelpers.ParseDDMMYYYY(s, MinYear, MaxYear)
		if !ok {
			status.SetText("Invalid date. Use DD.MM.YYYY")
			return
		}

		monthSelect.SetSelected(months[selectedMonth-1])
		yearEntry.SetText(strconv.Itoa(selectedYear))
		status.SetText("")
		render(selectedMonth, selectedYear)
	}

	// Initial render.
	render(selectedMonth, selectedYear)

	// Extend width of year entry.
	yearWrap := container.NewGridWrap(fyne.NewSize(80, yearEntry.MinSize().Height), yearEntry)

	// Join control elements.
	controls := container.NewHBox(
		widget.NewLabel("Month:"),
		monthSelect,
		widget.NewLabel("Year:"),
		yearWrap,
	)

	// Join all elements in the window.
	w.SetContent(container.NewVBox(
		// The client did not pay for an exit button.
		controls,
		jumpEntry,
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

	return container.NewStack(bg, container.NewCenter(t))
}

func colNormal() color.Color { return color.RGBA{255, 255, 255, 255} }

func colHoliday() color.Color { return color.RGBA{200, 255, 200, 255} }

func colSunday() color.Color { return color.RGBA{255, 200, 200, 255} }

func colHolidaySunday() color.Color { return color.RGBA{255, 240, 200, 255} }

func indexOf(list []string, value string) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}
