package ui

import (
	"log"
	"time"

	"gopad/internal/app"

	fyne "fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CustomEntry struct {
	widget.Entry
	OnShortcut func(shortcut fyne.Shortcut)
}

func NewCustomEntry() *CustomEntry {
	e := &CustomEntry{}
	e.MultiLine = true
	e.ExtendBaseWidget(e)
	return e
}

func (e *CustomEntry) TypedShortcut(shortcut fyne.Shortcut) {
	if e.OnShortcut != nil {
		e.OnShortcut(shortcut)
	}
	e.Entry.TypedShortcut(shortcut)
}

func NewWindow() {
	a := fyneApp.New()
	w := a.NewWindow("Gopad")

	svc := &app.Service{Window: w}

	editor := NewCustomEntry()
	status := widget.NewLabel("Ready")

	svc.TextContent = func() string { return editor.Text }
	svc.SetContent = func(text string) { editor.SetText(text) }
	svc.OnStatus = func(msg string) { status.SetText(msg) }

	// Shortcut handling inside the editor
	editor.OnShortcut = func(shortcut fyne.Shortcut) {
		switch s := shortcut.(type) {
		case *desktop.CustomShortcut:
			switch {
			case s.KeyName == fyne.KeyS && s.Modifier == fyne.KeyModifierControl:
				log.Println("Shortcut triggered: Ctrl+S")
				svc.Save()
			case s.KeyName == fyne.KeyO && s.Modifier == fyne.KeyModifierControl:
				log.Println("Shortcut triggered: Ctrl+O")
				svc.Open()
			case s.KeyName == fyne.KeyN && s.Modifier == fyne.KeyModifierControl:
				log.Println("Shortcut triggered: Ctrl+N")
				svc.New()
			}
		}
	}
	// Theme toggler
	themeSwitch := widget.NewCheck("", func(on bool) {
		if on {
			a.Settings().SetTheme(theme.DarkTheme())
			status.SetText("Switched to Dark Theme")
		} else {
			a.Settings().SetTheme(theme.LightTheme())
			status.SetText("Switched to Light Theme")
		}
	})
	themeToggle := container.NewHBox(widget.NewLabel("Dark Mode"), themeSwitch)

	// Also add global shortcuts to canvas for redundancy
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

	// Menus
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("New", svc.New),
		fyne.NewMenuItem("Open", svc.Open),
		fyne.NewMenuItem("Save", svc.Save),
		fyne.NewMenuItem("Save As", svc.SaveAs),
	)

	editMenu := fyne.NewMenu("Edit",
		fyne.NewMenuItem("Insert Date/Time", func() {
			editor.SetText(editor.Text + "\n" + time.Now().Format("2006-01-02 15:04:05"))
			status.SetText("Inserted date/time")
		}),
		fyne.NewMenuItem("Clear", func() {
			editor.SetText("")
			status.SetText("Cleared editor")
		}),
		fyne.NewMenuItem("Copy", func() {
			a.Clipboard().SetContent(editor.SelectedText())
			status.SetText("Copied text")
		}),
		fyne.NewMenuItem("Cut", func() {
			// Placeholder for cut logic
			status.SetText("Cut text")
		}),
		fyne.NewMenuItem("Paste", func() {
			content := a.Clipboard().Content()
			if content == "" {
				dialog.ShowInformation("Paste", "Clipboard is empty", w)
				return
			}
			editor.SetText(editor.Text + content)
			status.SetText("Pasted text")
		}),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowInformation("About", "A simple Notepad built with Fyne", w)
		}),
	)

	w.SetMainMenu(fyne.NewMainMenu(fileMenu, editMenu, helpMenu))

	statusBar := container.NewBorder(nil, nil, nil, status, themeToggle)
	content := container.NewBorder(nil, statusBar, nil, nil, editor)
	w.SetContent(content)

	w.Resize(fyne.NewSize(700, 500))
	w.ShowAndRun()
}
