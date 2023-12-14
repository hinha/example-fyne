package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/hinha/example-fyne/ui/menu"
)

type UI struct {
	EditWidget    *widget.Entry
	previewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem

	window fyne.Window
}

func (ui *UI) MakeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	ui.EditWidget = edit
	ui.previewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (ui *UI) CreateMenuItems(w fyne.Window) {
	m := menu.MainMenu{}
	ui.window = w

	fileMenu := m.CreateMenuFile(w, ui.EditWidget, ui.CurrentFile)
	viewMenu := ui.viewMainMenu()
	mainMenu := fyne.NewMainMenu(fileMenu, viewMenu)
	w.SetMainMenu(mainMenu)
}

func (ui *UI) viewMainMenu() *fyne.Menu {
	return fyne.NewMenu("View", fyne.NewMenuItem("Font", func() {

	}))
}
