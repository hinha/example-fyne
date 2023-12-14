package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io"
	"strings"
)

var Filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

type MainMenu struct {
	window fyne.Window

	saveMenuItem *fyne.MenuItem
	editWidget   *widget.Entry
	currentFile  fyne.URI
}

func (m *MainMenu) CreateMenuFile(window fyne.Window, editWidget *widget.Entry, currentFile fyne.URI) *fyne.Menu {
	// create open, save
	m.window = window
	m.editWidget = editWidget
	m.currentFile = currentFile
	openMenu := fyne.NewMenuItem("Open", m.actionMenuOpen)
	saveMenu := fyne.NewMenuItem("Save", m.actionSaveMenu)
	m.saveMenuItem = saveMenu
	m.saveMenuItem.Disabled = true

	saveAsMenu := fyne.NewMenuItem("Save as...", m.actionSaveAsMenu)

	return fyne.NewMenu("File", openMenu, saveMenu, saveAsMenu)
}

func (m *MainMenu) actionMenuOpen() {
	openDialog := dialog.NewFileOpen(func(readCloser fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, m.window)
			return
		}

		if readCloser == nil {
			return
		}
		defer readCloser.Close()

		data, err := io.ReadAll(readCloser)
		if err != nil {
			dialog.ShowError(err, m.window)
			return
		}

		m.editWidget.SetText(string(data))
		m.editWidget.SetText(string(data))
		m.currentFile = readCloser.URI()
		m.window.SetTitle(m.window.Title() + " - " + readCloser.URI().Name())
		m.saveMenuItem.Disabled = false
	}, m.window)
	openDialog.SetFilter(Filter)
	openDialog.Show()
}

func (m *MainMenu) actionSaveMenu() {
	if m.currentFile != nil {
		writer, err := storage.Writer(m.currentFile)
		if err != nil {
			dialog.ShowError(err, m.window)
		}

		writer.Write([]byte(m.editWidget.Text))
		defer writer.Close()

	}
}

func (m *MainMenu) actionSaveAsMenu() {
	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, m.window)
			return
		}

		if writer == nil {
			// cancel user
			return
		}

		if !strings.HasSuffix(strings.ToLower(writer.URI().String()), ".md") {
			dialog.ShowInformation("Error", "Please name your file with a .md extension!", m.window)
			return
		}

		//save file
		_, _ = writer.Write([]byte(m.editWidget.Text))

		m.currentFile = writer.URI()
		defer writer.Close()

		m.window.SetTitle(m.window.Title() + " - " + writer.URI().Name())
		m.saveMenuItem.Disabled = false
		return
	}, m.window)
	saveDialog.SetFileName("untitled.md")
	saveDialog.SetFilter(Filter)
	saveDialog.Show()
}
