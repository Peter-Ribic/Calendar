package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
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

	// 7 columns × 6 rows = 42 cells (always enough)
	cells := make([]fyne.CanvasObject, 0, 42)
	for i := 0; i < 42; i++ {
		bg := canvas.NewRectangle(color.RGBA{255, 255, 255, 255})
		bg.SetMinSize(fyne.NewSize(44, 34))
		txt := canvas.NewText("", color.Black)
		txt.Alignment = fyne.TextAlignCenter
		txt.TextSize = 14
		cells = append(cells, container.NewMax(bg, container.NewCenter(txt)))
	}

	grid := container.NewGridWithColumns(7, cells...)

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
