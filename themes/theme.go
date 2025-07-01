package themes

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CustomTheme struct {
	bg   color.Color
	fg   color.Color
	font fyne.Resource
}

func (m CustomTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		return m.bg
	case theme.ColorNameForeground:
		return m.fg
	default:
		return theme.DefaultTheme().Color(n, v)
	}
}

func (m CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func CustomThemeForm(app fyne.App) fyne.CanvasObject {
	bgEntry := widget.NewEntry()
	bgEntry.SetPlaceHolder("#RRGGBB (Background)")

	fgEntry := widget.NewEntry()
	fgEntry.SetPlaceHolder("#RRGGBB (Text)")

	form := widget.NewForm(
		widget.NewFormItem("Background Color", bgEntry),
		widget.NewFormItem("Font Color", fgEntry),
	)

	applyBtn := widget.NewButton("Apply", func() {
		bg, err1 := parseHexColor(bgEntry.Text)
		fg, err2 := parseHexColor(fgEntry.Text)
		if err1 == nil && err2 == nil {
			app.Settings().SetTheme(CustomTheme{bg: bg, fg: fg})
		}
	})

	return container.NewVBox(form, applyBtn)
}

func parseHexColor(s string) (color.Color, error) {
	var r, g, b uint8
	_, err := sscanf(s, "#%02x%02x%02x", &r, &g, &b)
	return color.RGBA{r, g, b, 0xff}, err
}

func sscanf(str string, format string, args ...interface{}) (int, error) {
	return fmt.Sscanf(str, format, args...)
}
