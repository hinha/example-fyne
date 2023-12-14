package action

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/hinha/example-fyne/ui"
	"io"
)

type Action struct {
	UI     *ui.UI
	Window fyne.Window
}

func (c *Action) OpenMenu() {
	openDialog := dialog.NewFileOpen(func(readCloser fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, c.Window)
			return
		}

		if readCloser == nil {
			return
		}
		defer readCloser.Close()

		data, err := io.ReadAll(readCloser)
		if err != nil {
			dialog.ShowError(err, c.Window)
			return
		}

		c.UI.EditWidget.SetText(string(data))
		c.UI.EditWidget.SetText(string(data))
		c.UI.CurrentFile = readCloser.URI()
		c.Window.SetTitle(c.Window.Title() + " - " + readCloser.URI().Name())
		//fmt.Println(m.saveItem)
		c.UI.SaveMenuItem.Disabled = false
	}, c.Window)
	openDialog.SetFilter(ui.Filter)
	openDialog.Show()
}
