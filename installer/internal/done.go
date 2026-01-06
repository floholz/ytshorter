package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type DoneStep struct {
	Stepper *TStepper
}

func (d DoneStep) Title() string {
	return "Done"
}

func (d DoneStep) Content() fyne.CanvasObject {
	return container.NewCenter(widget.NewLabel("All Done!"))
}

func (d DoneStep) OnInit() {
	if d.Stepper != nil && d.Stepper.Footer.Next != nil {
		d.Stepper.Footer.Next.Text = "Finish & Close"
	}
}

func (d DoneStep) OnNext() {
	fyne.CurrentApp().Quit()
}

func (d DoneStep) OnPrevious() {
	if d.Stepper != nil && d.Stepper.Footer.Next != nil {
		d.Stepper.Footer.Next.Text = "Next"
	}
}
