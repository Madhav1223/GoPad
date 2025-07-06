package statusbar

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2/widget"
)

func UpdateStatusBar(editor *widget.Entry, statusLabel *widget.Label) {
	text := editor.Text
	lines := len(strings.Split(text, "\n"))
	words := len(strings.Fields(text))
	characters := len(text)

	statusLabel.SetText(fmt.Sprintf("Line: %d | Words: %d | Characters: %d", lines, words, characters))
}
