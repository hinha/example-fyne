package main

import (
	"fyne.io/fyne/v2"
	application "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	out *widget.Label
}

var myApp App

func main() {
	app := application.New()
	w := app.NewWindow("hello")
	w.Resize(fyne.NewSize(300, 200))

	output, entry, btn := myApp.makeUI()

	w.SetContent(container.NewVBox(output, entry, btn))
	w.ShowAndRun()

	app.Quit()
}
