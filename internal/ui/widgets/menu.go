package widgets

import (
	"gopad/internal/app"
	"gopad/themes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
)

func MakeMenu(a fyne.App, w fyne.Window, svc *app.Service, tabs *container.AppTabs, setStatus func(string)) *fyne.MainMenu {
	file := fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() {
			editor, getText, setText := NewEditor(svc)
			tab := &app.TabContent{
				Title:   "Untitled",
				Content: getText,
				SetText: setText,
			}
			svc.Tabs = append(svc.Tabs, tab)
			tabs.Append(container.NewTabItem("Untitled", editor))
			setStatus("New tab created")
		}),
	)

	themeMenu := fyne.NewMenu("Theme",
		fyne.NewMenuItem("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		fyne.NewMenuItem("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
		fyne.NewMenuItem("Custom", func() {
			dialog.ShowCustom("Custom Theme", "Apply", themes.CustomThemeForm(a), w)
		}),
	)

	return fyne.NewMainMenu(file, themeMenu)
}
