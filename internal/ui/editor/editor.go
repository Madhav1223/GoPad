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
		return // Skip default if handled
	}
	e.Entry.TypedShortcut(shortcut)
}

func (e *CustomEntry) GetText() string {
	return e.Text
}

func (e *CustomEntry) SetText(text string) {
	e.Entry.SetText(text)
}

func (e *CustomEntry) SetOnChanged(handler func(string)) {
	e.Entry.OnChanged = handler
}

// for shortcuts handling
func (e *CustomEntry) SetOnShortcut(handler func(fyne.Shortcut)) {
	e.OnShortcut = handler
}
