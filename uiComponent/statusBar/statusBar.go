package statusbar

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateStatusBar(editor *widget.Entry) fyne.CanvasObject {

	statusLabel := widget.NewLabel("Line: 1 | Words: 0 | Characters: 0")

	UpdateStatusBar(editor, statusLabel)

	editor.OnChanged = func(s string) {
		UpdateStatusBar(editor, statusLabel)
	}

	statusContainer := container.NewHBox(
		widget.NewIcon(theme.AccountIcon()),
		statusLabel,
	)

	return statusContainer
}

func UpdateStatusBar(editor *widget.Entry, statusLabel *widget.Label) {
	text := editor.Text
	lines := len(strings.Split(text, "\n"))
	words := len(strings.Fields(text))
	characters := len(text)

	statusLabel.SetText(fmt.Sprintf("Line: %d | Words: %d | Characters: %d", lines, words, characters))
}
