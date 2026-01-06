package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ExtensionStep struct {
	Stepper *TStepper
}

func (e ExtensionStep) Title() string {
	return "Install Browser Extension"
}

func (e ExtensionStep) Content() fyne.CanvasObject {
	return container.NewCenter(widget.NewLabel("Extension Step"))
}

func (e ExtensionStep) OnInit() {
	// do nothing
}

func (e ExtensionStep) OnNext() {
	// do nothing
}

func (e ExtensionStep) OnPrevious() {
	// do nothing
}
