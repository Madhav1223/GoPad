package ui

import (
	"fmt"

	"gopad/helpers"

	statusbar "gopad/uiComponent/statusBar"

	menuBar "gopad/uiComponent/menu"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewWindow(file string) {

	a := app.New()
	w := a.NewWindow("Gopad")
	file_content, err := helpers.OpenFile(file)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", file, err)
		return
	}

	fileMenu := menuBar.CreateAllMenus(a)
	editor := widget.NewMultiLineEntry()
	editor.SetPlaceHolder("Type your text here...")
	editor.SetText(file_content)

	statusContainer := statusbar.CreateStatusBar(editor)

	content := container.NewBorder(nil, statusContainer, nil, nil, editor)

	w.SetMainMenu(fileMenu)
	w.SetContent(content)

	w.Resize(fyne.NewSize(800, 600))

	w.ShowAndRun()
}
