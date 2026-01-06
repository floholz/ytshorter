package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type IStep interface {
	Title() string
	Content() fyne.CanvasObject
	OnInit()
	OnNext()
	OnPrevious()
}

type TStepper struct {
	Header fyne.CanvasObject
	Body   fyne.CanvasObject
	Footer struct {
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

	stepper.steps = []IStep{
		InfoStep{Stepper: stepper},
		ExtensionStep{Stepper: stepper},
		ApplicationStep{Stepper: stepper},
		HostManifestStep{Stepper: stepper},
		DoneStep{Stepper: stepper},
	}
	stepper.Update(stepper.steps[stepper.StepIndex])
	stepper.steps[stepper.StepIndex].OnInit()
	return stepper
}

func (s *TStepper) Update(step IStep) {
	header := widget.NewLabel(step.Title())
	header.Alignment = fyne.TextAlignCenter
	header.TextStyle.Bold = true
	s.Header = header

	s.Body = step.Content()

	s.Content.Objects[0] = container.NewBorder(
		s.Header,
		container.NewHBox(layout.NewSpacer(), s.Footer.Previous, s.Footer.Next),
		nil, nil,
		s.Body,
	)
	s.Content.Refresh()
}

func (s *TStepper) Next() {
	s.steps[s.StepIndex].OnNext()
	if s.StepIndex == len(s.steps)-1 {
		return
	}
	s.StepIndex++
	s.steps[s.StepIndex].OnInit()
	s.Update(s.steps[s.StepIndex])
	s.Content.Refresh()
}

func (s *TStepper) Previous() {
	s.steps[s.StepIndex].OnPrevious()
	if s.StepIndex == 0 {
		return
	}
	s.StepIndex--
	s.steps[s.StepIndex].OnInit()
	s.Update(s.steps[s.StepIndex])
}
