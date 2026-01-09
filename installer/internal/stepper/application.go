package stepper

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ApplicationStep struct {
	Stepper *TStepper
}

func (a *ApplicationStep) Title() string {
	return "Install Native Application"
}

func (a *ApplicationStep) Content() fyne.CanvasObject {
	return container.NewCenter(widget.NewLabel("Application Step"))
}

func (a *ApplicationStep) OnInit() {
	// do nothing
	a.Stepper.Footer.Hint = widget.NewLabel("Click Next to verify application.")
}

func (a *ApplicationStep) OnNext() bool {
	// do nothing
	return true
}

func (a *ApplicationStep) OnPrevious() bool {
	// do nothing
	return true
}
