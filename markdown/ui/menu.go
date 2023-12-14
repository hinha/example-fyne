package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"io"
	"strings"
)

var Filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func (ui *UI) OpenMenu() *fyne.MenuItem {
	return fyne.NewMenuItem("Open...", func() {
		openDialog := dialog.NewFileOpen(func(readCloser fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, ui.window)
				return
			}

			if readCloser == nil {
				return
			}
			defer readCloser.Close()

			data, err := io.ReadAll(readCloser)
			if err != nil {
				dialog.ShowError(err, ui.window)
				return
			}

			ui.EditWidget.SetText(string(data))
			ui.EditWidget.SetText(string(data))
			ui.CurrentFile = readCloser.URI()
			ui.window.SetTitle(ui.window.Title() + " - " + readCloser.URI().Name())
			//fmt.Println(m.saveItem)
			ui.SaveMenuItem.Disabled = false
		}, ui.window)
		openDialog.SetFilter(Filter)
		openDialog.Show()
	})
}

func (ui *UI) SaveMenu() *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if ui.CurrentFile != nil {
			writer, err := storage.Writer(ui.CurrentFile)
			if err != nil {
				dialog.ShowError(err, ui.window)
			}

			writer.Write([]byte(ui.EditWidget.Text))
			defer writer.Close()

		}
	})
}

func (ui *UI) SaveAsMenu() *fyne.MenuItem {
	return fyne.NewMenuItem("Save as...", func() {
		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, ui.window)
				return
			}

			if writer == nil {
				// cancel user
				return
			}

			if !strings.HasSuffix(strings.ToLower(writer.URI().String()), ".md") {
				dialog.ShowInformation("Error", "Please name your file with a .md extension!", ui.window)
				return
			}

			//save file
			_, _ = writer.Write([]byte(ui.EditWidget.Text))

			ui.CurrentFile = writer.URI()
			defer writer.Close()

			ui.window.SetTitle(ui.window.Title() + " - " + writer.URI().Name())
			ui.SaveMenuItem.Disabled = false
			return
		}, ui.window)
		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(Filter)
		saveDialog.Show()
	})
}
