package helpers

import (
	"gopad/internal/app"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func HandleShortcut(shortcut fyne.Shortcut, svc *app.Service, showStatus func(string)) {
	switch s := shortcut.(type) {
	case *desktop.CustomShortcut:
		switch {
		case s.KeyName == fyne.KeyS && s.Modifier == fyne.KeyModifierControl:
			log.Println("Shortcut: Ctrl+S")
			svc.Save()
			showStatus("Saved")
		case s.KeyName == fyne.KeyO && s.Modifier == fyne.KeyModifierControl:
			log.Println("Shortcut: Ctrl+O")
			svc.Open()
			showStatus("Opened")
		case s.KeyName == fyne.KeyN && s.Modifier == fyne.KeyModifierControl:
			log.Println("Shortcut: Ctrl+N")
			svc.New()
			showStatus("New File")
		}
	}
}

// tabs.OnChanged = func(tab *container.TabItem) {
// 	for i, t := range tabs.Items {
// 		if t == tab {
// 			svc.ActiveTab = i
// 			break
// 		}
// 	}
// }
