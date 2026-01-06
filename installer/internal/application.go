package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ApplicationStep struct {
	Stepper *TStepper
}

func (a ApplicationStep) Title() string {
	return "Install Native Application"
}

func (a ApplicationStep) Content() fyne.CanvasObject {
	return container.NewCenter(widget.NewLabel("Application Step"))
}

func (a ApplicationStep) OnInit() {
	// do nothing
}

func (a ApplicationStep) OnNext() {
	// do nothing
}

func (a ApplicationStep) OnPrevious() {
	// do nothing
}
