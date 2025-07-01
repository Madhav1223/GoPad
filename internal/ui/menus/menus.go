package menus

import (
	"gopad/internal/app"
	"gopad/themes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"time"
)

func SetupMainMenu(app fyne.App, w fyne.Window, svc *app.Service, editor *widget.Entry, showStatus func(string)) *fyne.MainMenu {
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("New", svc.New),
		fyne.NewMenuItem("Open", svc.Open),
		fyne.NewMenuItem("Save", svc.Save),
		fyne.NewMenuItem("Save As", svc.SaveAs),
	)
	editMenu := fyne.NewMenu("Edit",
		fyne.NewMenuItem("Insert Date/Time", func() {
			editor.SetText(editor.Text + "\n" + time.Now().Format("2006-01-02 15:04:05"))
			showStatus("Inserted date/time")
		}),
		fyne.NewMenuItem("Clear", func() {
			editor.SetText("")
			showStatus("Cleared editor")
		}),
		fyne.NewMenuItem("Copy", func() {
			app.Clipboard().SetContent(editor.SelectedText())
			showStatus("Copied text")
		}),
		fyne.NewMenuItem("Paste", func() {
			content := app.Clipboard().Content()
			editor.SetText(editor.Text + content)
			showStatus("Pasted text")
		}),
	)

	themeMenu := fyne.NewMenu("Themes",
		fyne.NewMenuItem("Dark Mode", func() {
			app.Settings().SetTheme(theme.DarkTheme())
			showStatus("Switched to Dark Theme")
		}),
		fyne.NewMenuItem("Light Mode", func() {
			app.Settings().SetTheme(theme.LightTheme())
			showStatus("Switched to Light Theme")
		}),
		fyne.NewMenuItem("Custom Theme", func() {
			dialog.ShowCustom("Create Theme", "Apply", themes.CustomThemeForm(app), w)
		}),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowInformation("About", "A simple Notepad built with Fyne", w)
		}),
	)

	return fyne.NewMainMenu(fileMenu, editMenu, themeMenu, helpMenu)
}
