package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type HostManifestStep struct {
	Stepper *TStepper
}

func (h HostManifestStep) Title() string {
	return "Setup Native Host Manifest"
}

func (h HostManifestStep) Content() fyne.CanvasObject {
	return container.NewCenter(widget.NewLabel("Host Manifest Step"))
}

func (h HostManifestStep) OnInit() {
	// do nothing
}

func (h HostManifestStep) OnNext() {
	// do nothing
}

func (h HostManifestStep) OnPrevious() {
	// do nothing
}
