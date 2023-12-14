package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"github.com/hinha/example-fyne/ui/menu"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_makeUI(t *testing.T) {
	var testUI UI

	edit, preview := testUI.MakeUI()

	test.Type(edit, "Edit")
	if preview.String() != "Edit" {
		t.Fail()
	}
}

func Test_RunApp(t *testing.T) {
	var testUI UI

	testApp := test.NewApp()
	testWindow := testApp.NewWindow("Test MarkDown")

	edit, preview := testUI.MakeUI()
	testUI.CreateMenuItems(testWindow)
	testWindow.SetContent(container.NewHSplit(edit, preview))

	testApp.Run()
	test.Type(edit, "Some Text")
	if preview.String() != "Some Text" {
		t.Fail()
	}
}

func Test_MainMenuItems(t *testing.T) {
	var testUI UI
	var testMenu menu.MainMenu

	testApp := test.NewApp()
	testWindow := testApp.NewWindow("Test MarkDown")

	testUI.CreateMenuItems(testWindow)
	assert.Equal(t, len(testWindow.MainMenu().Items), 2)

	mainMenuFile := testMenu.CreateMenuFile(testWindow, testUI.EditWidget, testUI.CurrentFile)
	mainMenuView := testUI.viewMainMenu()

	testMainMenu := testWindow.MainMenu()
	assert.Equal(t, testMainMenu.Items[0].Label, mainMenuFile.Label)
	assert.Equal(t, testMainMenu.Items[1].Label, mainMenuView.Label)
}

func Test_MenuItemFile(t *testing.T) {
	var testUI UI
	var testMenu menu.MainMenu

	testApp := test.NewApp()
	testWindow := testApp.NewWindow("Test MarkDown")

	testUI.CreateMenuItems(testWindow)
	assert.Equal(t, len(testWindow.MainMenu().Items), 2)

	mainMenuFile := testMenu.CreateMenuFile(testWindow, testUI.EditWidget, testUI.CurrentFile)

	testMainMenu := testWindow.MainMenu()

	// Test open menu
	assert.Equal(t, testMainMenu.Items[0].Label, mainMenuFile.Label)
	assert.Equal(t, testMainMenu.Items[0].Items[0].Label, mainMenuFile.Items[0].Label) // label: open
	assert.Equal(t, testMainMenu.Items[0].Items[1].Label, mainMenuFile.Items[1].Label) // label: save
	assert.Equal(t, testMainMenu.Items[0].Items[2].Label, mainMenuFile.Items[2].Label) // label: save as
}
