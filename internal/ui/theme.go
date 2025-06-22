package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type forcedTheme struct {
	base    fyne.Theme
	variant fyne.ThemeVariant
}

func (f forcedTheme) Color(n fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return f.base.Color(n, f.variant)
}

func (f forcedTheme) Font(s fyne.TextStyle) fyne.Resource {
	return f.base.Font(s)
}

func (f forcedTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return f.base.Icon(n)
}

func (f forcedTheme) Size(n fyne.ThemeSizeName) float32 {
	return f.base.Size(n)
}

func (f forcedTheme) Variant() fyne.ThemeVariant {
	return f.variant
}
