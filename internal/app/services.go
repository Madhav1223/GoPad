package app

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

type Service struct {
	CurrentFile fyne.URI
	Window      fyne.Window

	TextContent func() string
	SetContent  func(string)
	OnStatus    func(string)
}

func (s *Service) Save() {
	if s.CurrentFile == nil {
		s.SaveAs()
		return
	}
	writer, err := storage.Writer(s.CurrentFile)
	if err != nil {
		dialog.ShowError(err, s.Window)
		return
	}
	defer writer.Close()

	_, err = io.WriteString(writer, s.TextContent())
	if err != nil {
		dialog.ShowError(err, s.Window)
		return
	}

	s.OnStatus("Saved: " + s.CurrentFile.Name())
}

func (s *Service) SaveAs() {
	dialog.ShowFileSave(func(wc fyne.URIWriteCloser, err error) {
		if err != nil || wc == nil {
			s.OnStatus("Save canceled")
			return
		}
		defer wc.Close()

		_, err = io.WriteString(wc, s.TextContent())
		if err != nil {
			dialog.ShowError(err, s.Window)
			return
		}

		s.CurrentFile = wc.URI()
		s.OnStatus("Saved as: " + s.CurrentFile.Name())
	}, s.Window)
}

func (s *Service) Open() {
	dialog.ShowFileOpen(func(rc fyne.URIReadCloser, err error) {
		if err != nil || rc == nil {
			s.OnStatus("Open canceled")
			return
		}
		defer rc.Close()

		data, err := io.ReadAll(rc)
		if err != nil {
			dialog.ShowError(err, s.Window)
			return
		}

		s.CurrentFile = rc.URI()
		s.SetContent(string(data))
		s.OnStatus("Opened: " + s.CurrentFile.Name())
	}, s.Window)
}

func (s *Service) New() {
	s.CurrentFile = nil
	s.SetContent("")
	s.OnStatus("New file created")
}
