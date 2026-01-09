package stepper

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type HostManifestStep struct {
	Stepper *TStepper
}

func (h *HostManifestStep) Title() string {
	return "Setup Native Host Manifest"
}

func (h *HostManifestStep) Content() fyne.CanvasObject {
	return container.NewCenter(widget.NewLabel("Host Manifest Step"))
}

func (h *HostManifestStep) OnInit() {
	h.Stepper.Footer.Hint = widget.NewLabel("Click Next to setup host manifest.")
}

func (h *HostManifestStep) OnNext() bool {
	// do nothing
	return true
}

func (h *HostManifestStep) OnPrevious() bool {
	// do nothing
	return true
}
