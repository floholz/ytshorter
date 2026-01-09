package stepper

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type InfoStep struct {
	Stepper *TStepper
}

func (i *InfoStep) Title() string {
	return "Install Process"
}

func (i *InfoStep) Content() fyne.CanvasObject {
	return container.NewCenter(
		container.NewVBox(
			widget.NewLabel("This installer will lead you through the process of"),
			widget.NewForm(
				widget.NewFormItem("1.", widget.NewLabel("Loading the browser extension")),
				widget.NewFormItem("2.", widget.NewLabel("Installing the native application")),
				widget.NewFormItem("3.", widget.NewLabel("Setting up the host manifest")),
			),
		),
	)
}

func (i *InfoStep) OnInit() {
	// do nothing
	if i.Stepper != nil && i.Stepper.Footer.Previous != nil {
		i.Stepper.Footer.Previous.Disable()
		i.Stepper.Footer.Previous.Hide()
	}
	i.Stepper.Footer.Hint = widget.NewLabel("Press 'Next' to proceed!")
}

func (i *InfoStep) OnNext() bool {
	// do nothing
	i.Stepper.Footer.Previous.Enable()
	i.Stepper.Footer.Previous.Show()
	return true
}

func (i *InfoStep) OnPrevious() bool {
	// do nothing
	return true
}
