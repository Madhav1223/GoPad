package ui

import (
	"fmt"

	"gopad/helpers"

	statusbar "gopad/uiComponent/statusBar"

	menuBar "gopad/uiComponent/menu"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
	fmt.Println("File content:")
	fmt.Println(file_content)

	fileMenu := menuBar.CreateAllMenus(a)
	editor := widget.NewMultiLineEntry()
	editor.SetPlaceHolder("Type your text here...")
	editor.SetText(file_content)

	statusLabel := widget.NewLabel("Line: 1 | Words: 0 | Characters: 0")
	statusbar.UpdateStatusBar(editor, statusLabel)
	editor.OnChanged = func(s string) {
		statusbar.UpdateStatusBar(editor, statusLabel)
	}

	statusContainer := container.NewHBox(widget.NewIcon(theme.AccountIcon()), statusLabel)

	content := container.NewBorder(nil, statusContainer, nil, nil, editor)

	w.SetMainMenu(fileMenu)
	w.SetContent(content)

	w.Resize(fyne.NewSize(800, 600))

	w.ShowAndRun()
}
