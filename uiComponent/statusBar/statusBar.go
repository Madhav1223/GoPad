package statusbar

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2/widget"
)

//  Initalize statusbar taking filecontent as input
//  and update the statusbar with line, word, and character count

func InitializeStatusBar(editor *widget.Entry) *widget.Label {
	statusLabel := widget.NewLabel("Line: 1 | Words: 0 | Characters: 0")
	UpdateStatusBar(editor, statusLabel)

	editor.OnChanged = func(s string) {
		UpdateStatusBar(editor, statusLabel)
	}

	return statusLabel
}
func UpdateStatusBar(editor *widget.Entry, statusLabel *widget.Label) {
	text := editor.Text
	lines := len(strings.Split(text, "\n"))
	words := len(strings.Fields(text))
	characters := len(text)

	statusLabel.SetText(fmt.Sprintf("Line: %d | Words: %d | Characters: %d", lines, words, characters))
}
