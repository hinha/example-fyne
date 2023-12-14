package main

import "fyne.io/fyne/v2/widget"

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello World")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		app.out.SetText(entry.Text)
	})
	btn.Importance = widget.HighImportance
	app.out = output

	return output, entry, btn
}
