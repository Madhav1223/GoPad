package ui

import (
	"gopad/internal/app"
	"gopad/internal/ui/editor"
	editorHelpers "gopad/internal/ui/editor/helpers"
	"gopad/internal/ui/menus"
	"gopad/internal/ui/statusbar"
	"time"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	// Importing desktop package for custom shortcuts
)

func NewWindow(files []string) {
	a := fyneApp.New()
	w := a.NewWindow("Gopad")
	svc := &app.Service{Window: w}

	if len(files) > 0 {
		w.SetTitle("Gopad - " + files[0])
	} else {
		w.SetTitle("Gopad - Untitled")
	}

	// Initialize editor
	editorWidget := editor.NewCustomEntry(svc, nil)

	// Setup StatusBar with live stats function
	statusBar, statusBarUI := statusbar.NewStatusBar(func() string {
		return editorWidget.Text
	}, 3*time.Second) // Or replace with helpers.StatusResetDelay if you have one

	// Link status updates on change
	editorWidget.OnChanged = func(_ string) {
		statusBar.UpdateLive()
	}

	// Handle keyboard shortcuts
	editorWidget.OnShortcut = func(sc fyne.Shortcut) {
		editorHelpers.HandleShortcut(sc, svc, statusBar.ShowMessage)
	}

	// Hook into service
	svc.TextContent()
	svc.SetContent("Test")
	svc.OnStatus = statusBar.ShowMessage

	// Setup menu
	mainMenu := menus.SetupMainMenu(a, w, svc, &editorWidget.Entry, statusBar.ShowMessage)
	w.SetMainMenu(mainMenu)

	// Keyboard shortcuts
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

	// Layout: editor as center, status bar as bottom
	content := container.NewBorder(nil, statusBarUI, nil, nil, editorWidget)
	w.SetContent(content)
	w.Resize(fyne.NewSize(700, 500))
	w.ShowAndRun()
}
