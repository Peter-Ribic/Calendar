package main

import (
	"fyne.io/fyne/v2/app"

	"github.com/Peter-Ribic/Calendar/internal/ui"
)

func main() {
	a := app.New()
	ui.Run(a)
}
