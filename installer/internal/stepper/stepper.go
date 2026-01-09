package stepper

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type IStep interface {
	Title() string
	Content() fyne.CanvasObject
	OnInit()
	OnNext() bool
	OnPrevious() bool
}

type TStepper struct {
	Header fyne.CanvasObject
	Body   fyne.CanvasObject
	Footer struct {
		Hint           fyne.CanvasObject
		Previous, Next *widget.Button
	}
	Content   *fyne.Container
	StepIndex int
	steps     []IStep
}

func NewStepper() *TStepper {
	stepper := &TStepper{
		StepIndex: 0,
		Content:   container.NewPadded(widget.NewLabel("loading...")),
	}
	stepper.Footer.Previous = widget.NewButton("Previous", stepper.Previous)
	stepper.Footer.Next = widget.NewButton("Next", stepper.Next)
	stepper.Footer.Hint = widget.NewLabel("Press 'Next' to proceed!")

	stepper.steps = []IStep{
		&InfoStep{Stepper: stepper},
		&ExtensionStep{Stepper: stepper},
		&ApplicationStep{Stepper: stepper},
		&HostManifestStep{Stepper: stepper},
		&DoneStep{Stepper: stepper},
	}
	stepper.Update(stepper.steps[stepper.StepIndex])

	stepper.steps[stepper.StepIndex].OnInit()
	return stepper
}

func (s *TStepper) Update(step IStep) {
	headerLabel := widget.NewLabel(step.Title())
	headerLabel.Alignment = fyne.TextAlignCenter
	headerLabel.TextStyle.Bold = true

	progress := widget.NewProgressBar()
	progress.TextFormatter = func() string {
		return "[ " + strconv.Itoa(s.StepIndex+1) + " / " + strconv.Itoa(len(s.steps)) + " ]"
	}
	progress.Max = float64(len(s.steps) - 1)
	progress.Value = float64(s.StepIndex)

	s.Header = container.NewVBox(progress, headerLabel, widget.NewSeparator())

	s.Body = step.Content()

	s.Content.Objects[0] = container.NewBorder(
		s.Header,
		container.NewHBox(s.Footer.Hint, layout.NewSpacer(), s.Footer.Previous, s.Footer.Next),
		nil, nil,
		s.Body,
	)
	s.Content.Refresh()
}

func (s *TStepper) Next() {
	if !s.steps[s.StepIndex].OnNext() {
		s.Update(s.steps[s.StepIndex])
		return
	}
	if s.StepIndex == len(s.steps)-1 {
		return
	}
	s.StepIndex++
	s.steps[s.StepIndex].OnInit()
	s.Update(s.steps[s.StepIndex])
	s.Content.Refresh()
}

func (s *TStepper) Previous() {
	if !s.steps[s.StepIndex].OnPrevious() {
		s.Update(s.steps[s.StepIndex])
		return
	}
	if s.StepIndex == 0 {
		return
	}
	s.StepIndex--
	s.steps[s.StepIndex].OnInit()
	s.Update(s.steps[s.StepIndex])
}
