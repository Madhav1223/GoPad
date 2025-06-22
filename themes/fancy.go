package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type FancyTheme struct{}

func (t FancyTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{255, 240, 245, 255} // light pink background
	case theme.ColorNameButton:
		return color.RGBA{255, 105, 180, 255} // hot pink buttons
	case theme.ColorNameDisabledButton:
		return color.RGBA{255, 182, 193, 255} // light pink disabled
	case theme.ColorNameDisabled:
		return color.RGBA{200, 200, 200, 255}
	case theme.ColorNamePlaceHolder:
		return color.RGBA{218, 112, 214, 255} // orchid placeholder
	case theme.ColorNameHover:
		return color.RGBA{255, 20, 147, 255} // deeper pink hover
	case theme.ColorNameFocus:
		return color.RGBA{138, 43, 226, 255} // blueviolet focus
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t FancyTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t FancyTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
