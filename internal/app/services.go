package app

import "fyne.io/fyne/v2"

type Service struct {
	Window      fyne.Window
	TextContent func() string
	SetContent  func(string)
	OnStatus    func(string)
}

func (s *Service) Save()    { s.OnStatus("Saved") }
func (s *Service) Open()    { s.OnStatus("Opened") }
func (s *Service) New()     { s.OnStatus("New File") }
func (s *Service) SaveAs()  { s.OnStatus("Save As not implemented") }
