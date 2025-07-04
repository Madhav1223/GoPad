// package app

// import "fyne.io/fyne/v2"

// type Service struct {
// 	Window      fyne.Window
// 	TextContent func() string
// 	SetContent  func(string)
// 	OnStatus    func(string)
// }

package app

import "fyne.io/fyne/v2"

type TabContent struct {
	Title   string
	GetText func() string
	SetText func(string)
	// Optional: Add Filename, IsModified, etc. if needed
}

type Service struct {
	Window    fyne.Window
	Tabs      []*TabContent
	ActiveTab int
	OnStatus  func(string)
}

// CurrentTab returns the currently active tab.
func (s *Service) CurrentTab() *TabContent {
	if s.ActiveTab >= 0 && s.ActiveTab < len(s.Tabs) {
		return s.Tabs[s.ActiveTab]
	}
	return nil
}

// TextContent gets the text of the current tab.
func (s *Service) TextContent() string {
	tab := s.CurrentTab()
	if tab != nil && tab.GetText != nil {
		return tab.GetText()
	}
	return ""
}

// SetContent sets the text of the current tab.
func (s *Service) SetContent(text string) {
	tab := s.CurrentTab()
	if tab != nil && tab.SetText != nil {
		tab.SetText(text)
	}
}

// Save shows a status message for saving (placeholder for actual save logic).
func (s *Service) Save() {
	tab := s.CurrentTab()
	if tab == nil {
		s.OnStatus("No active tab to save")
		return
	}
	// TODO: Add real file save logic here
	s.OnStatus("Saved: " + tab.Title)
}

// Open shows a status message (placeholder for file open logic).
func (s *Service) Open() {
	s.OnStatus("Opened file (not yet implemented)")
}

// New creates a new tab (placeholder for real tab creation).
func (s *Service) New() {
	s.OnStatus("New file (not yet implemented)")
}

// SaveAs shows a status message for Save As (not implemented yet).
func (s *Service) SaveAs() {
	s.OnStatus("Save As not yet implemented")
}
