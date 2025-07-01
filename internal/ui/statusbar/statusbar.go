package statusbar

import (
	"fmt"
	"time"

	"gopad/internal/ui/statusbar/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type StatusBar struct {
	label     *widget.Label
	getText   func() string
	resetTime time.Duration
}

func NewStatusBar(getText func() string, resetTime time.Duration) (*StatusBar, fyne.CanvasObject) {
	label := widget.NewLabel("")
	sb := &StatusBar{
		label:     label,
		getText:   getText,
		resetTime: resetTime,
	}
	sb.updateDefault()
	return sb, container.NewBorder(nil, nil, nil, label, widget.NewLabel("Gopad v1.0"))
}

func (s *StatusBar) ShowMessage(msg string) {
	s.label.SetText(msg)
	go func() {
		time.Sleep(s.resetTime)
		s.updateDefault()
	}()
}

func (s *StatusBar) UpdateLive() {
	s.updateDefault()
}

func (s *StatusBar) updateDefault() {
	stats := helpers.ComputeStats(s.getText())
	s.label.SetText(fmt.Sprintf("Lines: %d | Words: %d | Chars: %d", stats.Lines, stats.Words, stats.Chars))
}
