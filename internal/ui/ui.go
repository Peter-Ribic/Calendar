package ui

import (
	"image/color"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// Run creates the window and renders the current month into a manual grid.
func Run(a fyne.App) {
	w := a.NewWindow("Calendar")

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
			cellTexts[idx].Refresh()
		}
	}

	now := time.Now()
	render(int(now.Month()), now.Year())

	w.SetContent(container.NewVBox(
		weekdayHeader,
		grid,
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

func colNormal() color.Color {
	return color.RGBA{255, 255, 255, 255}
}

func daysIn(month int, year int) int {
	t := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.Local)
	return t.Day()
}

// Monday=0 ... Sunday=6
func weekdayMondayIndex(wd time.Weekday) int {
	return (int(wd) + 6) % 7
}
