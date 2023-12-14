package main

import (
	"fyne.io/fyne/v2"
	application "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/hinha/example-fyne/ui"
)

var cfg ui.UI

func main() {
	app := application.New()

	window := app.NewWindow("Markdown")

	edit, preview := cfg.MakeUI()
	cfg.CreateMenuItems(window)
	window.SetContent(container.NewHSplit(edit, preview))

	window.Resize(fyne.NewSize(800, 500))
	window.ShowAndRun()
	window.CenterOnScreen()
}
