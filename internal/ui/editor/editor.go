package editor

import (
	"gopad/internal/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type CustomEntry struct {
	widget.Entry
	OnShortcut func(fyne.Shortcut)
}

func NewCustomEntry(svc *app.Service, onShortcut func(fyne.Shortcut)) *CustomEntry {
	e := &CustomEntry{}
	e.MultiLine = true
	e.ExtendBaseWidget(e)
	e.OnShortcut = onShortcut
	return e
}

func (e *CustomEntry) TypedShortcut(shortcut fyne.Shortcut) {
	if e.OnShortcut != nil {
		e.OnShortcut(shortcut)
	}
	e.Entry.TypedShortcut(shortcut)
}
