package stepper

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type DoneStep struct {
	Stepper *TStepper
}

func (d *DoneStep) Title() string {
	return "All Done!"
}

func (d *DoneStep) Content() fyne.CanvasObject {
	return container.NewHBox(
		layout.NewSpacer(),
		container.NewVBox(
			layout.NewSpacer(),
			widget.NewLabel("YTShorter should be fully setup."),
			widget.NewLabel("You can now close this window."),
			layout.NewSpacer(),
			widget.NewLabelWithStyle("It's recommended to restart your browser now.", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: true}),
			layout.NewSpacer(),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)
}

func (d *DoneStep) OnInit() {
	if d.Stepper != nil && d.Stepper.Footer.Next != nil {
		d.Stepper.Footer.Next.Text = "Finish & Close"
	}
	d.Stepper.Footer.Hint = widget.NewLabel("Press 'Finish & Close' to close the installer.")
}

func (d *DoneStep) OnNext() bool {
	fyne.CurrentApp().Quit()
	return true
}

func (d *DoneStep) OnPrevious() bool {
	if d.Stepper != nil && d.Stepper.Footer.Next != nil {
		d.Stepper.Footer.Next.Text = "Next"
	}
	return true
}
