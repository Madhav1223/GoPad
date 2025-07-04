package widgets

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewStatusBar() (*fyne.Container, func(string)) {
	status := widget.NewLabel("Ready")
	bar := container.NewHBox(widget.NewLabel("Gopad"), status)
	return bar, func(msg string) {
		status.SetText(msg)
		go func() {
			time.Sleep(3 * time.Second)
			status.SetText("Ready")
		}()
	}
}

func UpdateLive