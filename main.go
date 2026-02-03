package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calendar – Fyne PoC")

	w.SetContent(container.NewVBox(
		widget.NewLabel("Fyne is working"),
		widget.NewButton("Close", func() { w.Close() }),
	))

	w.Resize(fyne.NewSize(300, 120))
	w.ShowAndRun()
}
