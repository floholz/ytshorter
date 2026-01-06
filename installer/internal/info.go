package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type InfoStep struct {
	Stepper *TStepper
}

func (i InfoStep) Title() string {
	return "YTShorter Installer"
}

func (i InfoStep) Content() fyne.CanvasObject {
	return container.NewCenter(widget.NewLabel("Info Step"))
}

func (i InfoStep) OnInit() {
	// do nothing
	if i.Stepper != nil && i.Stepper.Footer.Previous != nil {
		i.Stepper.Footer.Previous.Disable()
		i.Stepper.Footer.Previous.Hide()
	}
}

func (i InfoStep) OnNext() {
	// do nothing
	i.Stepper.Footer.Previous.Enable()
	i.Stepper.Footer.Previous.Show()
}

func (i InfoStep) OnPrevious() {
	// do nothing
}
