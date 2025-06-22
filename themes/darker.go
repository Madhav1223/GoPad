package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type DarkerProTheme struct{}

func (t DarkerProTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{18, 18, 18, 255} // almost black background
	case theme.ColorNameButton:
		return color.RGBA{0, 122, 204, 255} // deep blue buttons
	case theme.ColorNameDisabledButton:
		return color.RGBA{40, 40, 40, 255}
	case theme.ColorNameDisabled:
		return color.RGBA{100, 100, 100, 255}
	case theme.ColorNamePlaceHolder:
		return color.RGBA{130, 130, 130, 255}
	case theme.ColorNameHover:
		return color.RGBA{50, 50, 50, 255}
	case theme.ColorNameFocus:
		return color.RGBA{0, 153, 255, 255}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t DarkerProTheme) Font(style fyne.TextStyle) fyne.Resource {
	// Use default fonts for now, you can customize fonts here
	return theme.DefaultTheme().Font(style)
}

func (t DarkerProTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
