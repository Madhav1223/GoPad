package widgets

import (
	"fmt"
	"strings"

	"gopad/internal/app"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type CustomEntry struct {
	widget.Entry
	OnShortcut func(fyne.Shortcut)
}

func (e *CustomEntry) TypedShortcut(shortcut fyne.Shortcut) {
	if e.OnShortcut != nil {
		e.OnShortcut(shortcut)
	}
	e.Entry.TypedShortcut(shortcut)
}

func NewEditor(svc *app.Service) (*CustomEntry, func() string, func(string)) {
	entry := &CustomEntry{}
	entry.MultiLine = true
	entry.ExtendBaseWidget(entry)

	entry.OnChanged = func(s string) {
		if svc.OnStatus != nil {
			wordCount := len(strings.Fields(s))
			lineCount := len(strings.Split(s, "\n"))
			svc.OnStatus("Lines: " + itoa(lineCount) + " | Words: " + itoa(wordCount))
		}
	}

	entry.OnShortcut = func(s fyne.Shortcut) {
		if key, ok := s.(*desktop.CustomShortcut); ok {
			switch {
			case key.KeyName == fyne.KeyS && key.Modifier == fyne.KeyModifierControl:
				svc.OnStatus("Saved (mock)")
			}
		}
	}

	return entry, func() string {
			return entry.Text
		}, func(t string) {
			entry.SetText(t)
		}
}

func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}
