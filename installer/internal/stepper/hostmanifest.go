package stepper

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/floholz/ytshorter/installer/internal/utils"
)

type HostManifestStep struct {
	Stepper        *TStepper
	error          error
	errorDismissed bool
}

func (h *HostManifestStep) Title() string {
	return "Setup Native Host Manifest"
}

func (h *HostManifestStep) Content() fyne.CanvasObject {
	if h.error != nil {
		return container.NewCenter(
			container.NewVBox(
				widget.NewLabel("Error creating the native host manifest:"),
				widget.NewLabel(h.error.Error()),
			),
		)
	}

	manifestPath, err := utils.GetNativeHostManifestPath()
	if err != nil {
		manifestPath = "???"
	}

	copyHint := widget.NewLabelWithStyle("(click to copy)", fyne.TextAlignTrailing, fyne.TextStyle{Italic: true})

	hyLink := widget.NewHyperlink(utils.TruncatePath(manifestPath, 40), nil)
	hyLink.OnTapped = func() {
		fyne.CurrentApp().Clipboard().SetContent(manifestPath)
		copyHint.Text = "Copied!"
		copyHint.Refresh()

		go func() {
			waiting := 5
			ticker := time.NewTicker(time.Second)
			for range ticker.C {
				fyne.Do(func() {
					copyHint.Text = "Copied! (" + strconv.Itoa(waiting) + ")"
					copyHint.Refresh()
					waiting--
					if waiting < 0 {
						ticker.Stop()
						copyHint.Text = "(click to copy)"
						copyHint.Refresh()
					}
				})
			}
		}()
	}

	return container.NewVBox(
		layout.NewSpacer(),
		widget.NewLabel("The Native Host Manifest is being created at:"),
		container.NewHBox(
			hyLink,
			copyHint,
		),
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
	)
}

func (h *HostManifestStep) OnInit() {
	h.Stepper.Footer.Hint = widget.NewLabel("Click Next to setup host manifest.")

	if h.Stepper != nil && h.Stepper.Footer.Next != nil {
		h.Stepper.Footer.Next.Text = "Next"
	}

	h.error = utils.InstallManifest()
	if h.error != nil {
		fmt.Println(h.error)
	}
}

func (h *HostManifestStep) OnNext() bool {
	if h.error != nil {
		if !h.errorDismissed {
			h.errorDismissed = true

			if h.Stepper != nil && h.Stepper.Footer.Next != nil {
				h.Stepper.Footer.Next.Text = "Ignore"
			}
			return false
		}
		return true
	}

	return true
}

func (h *HostManifestStep) OnPrevious() bool {
	if h.Stepper != nil && h.Stepper.Footer.Next != nil {
		h.Stepper.Footer.Next.Text = "Next"
	}
	return true
}
