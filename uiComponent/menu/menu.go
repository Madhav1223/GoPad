package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func CreateAllMenus(a fyne.App) *fyne.MainMenu {
	fileMenu := createFileMenu(a)
	editMenu := createEditMenu(a)
	themeMenu := createThemeMenu(a)
	helpMenu := createHelpMenu(a)

	return fyne.NewMainMenu(fileMenu, editMenu, themeMenu, helpMenu)
}

func createFileMenu(a fyne.App) *fyne.Menu {
	return fyne.NewMenu("File",
		fyne.NewMenuItem("Open", func() {

			println("Open File clicked")
		}),
		fyne.NewMenuItem("Save", func() {

			println("Save File clicked")
		}),
		fyne.NewMenuItem("Quit", func() {

			println("Quit clicked")
			a.Quit()
		}),
	)
}

func createEditMenu(a fyne.App) *fyne.Menu {
	return fyne.NewMenu("Edit",
		fyne.NewMenuItem("Insert Date/Time", func() {

			println("Insert Date/Time clicked")
		}),
		fyne.NewMenuItem("Clear", func() {

			println("Clear clicked")
		}),
		fyne.NewMenuItem("Copy", func() {

			println("Copy clicked")
		}),
		fyne.NewMenuItem("Paste", func() {

			println("Paste clicked")
		}),
	)
}

func createThemeMenu(a fyne.App) *fyne.Menu {
	return fyne.NewMenu("Themes",
		fyne.NewMenuItem("Dark Mode", func() {

			println("Dark Mode clicked")
		}),
		fyne.NewMenuItem("Light Mode", func() {

			println("Light Mode clicked")
		}),
		fyne.NewMenuItem("Custom Theme", func() {

			println("Custom Theme clicked")
		}),
	)
}

func createHelpMenu(a fyne.App) *fyne.Menu {
	return fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {

			dialog.ShowInformation("About", "This is a simple notepad app built with Fyne.", nil)
		}),
	)
}
