package ui

import (
	"gopad/internal/app"
	"gopad/internal/ui/editor"
	editorHelpers "gopad/internal/ui/editor/helpers"
	"gopad/internal/ui/menus"
	"gopad/internal/ui/statusbar"
	"gopad/internal/ui/statusbar/helpers"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
)

func NewWindow() {
	a := fyneApp.New()
	w := a.NewWindow("Gopad")

	svc := &app.Service{Window: w}
	editorWidget := editor.NewCustomEntry(svc, nil)

	statusBar, statusBarUI := statusbar.NewStatusBar(func() string {
		return editorWidget.Text
	}, helpers.StatusResetDelay)

	editorWidget.OnChanged = func(_ string) {
		statusBar.UpdateLive()
	}
	editorWidget.OnShortcut = func(sc fyne.Shortcut) {
		editorHelpers.HandleShortcut(sc, svc, statusBar.ShowMessage)
	}

	svc.TextContent = func() string { return editorWidget.Text }
	svc.SetContent = func(text string) { editorWidget.SetText(text) }
	svc.OnStatus = statusBar.ShowMessage

	mainMenu := menus.SetupMainMenu(a, w, svc, &editorWidget.Entry, statusBar.ShowMessage)
	w.SetMainMenu(mainMenu)

	canvas := w.Canvas()
	canvas.AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: fyne.KeyModifierControl}, func(fyne.Shortcut) {
		svc.Save()
	})
	canvas.AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyO, Modifier: fyne.KeyModifierControl}, func(fyne.Shortcut) {
		svc.Open()
	})
	canvas.AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyN, Modifier: fyne.KeyModifierControl}, func(fyne.Shortcut) {
		svc.New()
	})

	content := container.NewBorder(nil, statusBarUI, nil, nil, editorWidget)
	w.SetContent(content)
	w.Resize(fyne.NewSize(700, 500))
	w.ShowAndRun()
}
