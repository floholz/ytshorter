package stepper

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/floholz/ytshorter/installer/internal/goapp"
	"github.com/floholz/ytshorter/installer/internal/utils"
)

type ApplicationStep struct {
	Stepper *TStepper
}

func (a *ApplicationStep) Title() string {
	return "Install Native Application"
}

func (a *ApplicationStep) Content() fyne.CanvasObject {
	appPath, err := utils.GetAppPath()
	if err != nil {
		appPath = "???"
	}

	copyHint := widget.NewLabelWithStyle("(click to copy)", fyne.TextAlignTrailing, fyne.TextStyle{Italic: true})

	appPathHyLink := widget.NewHyperlink(appPath, nil)
	appPathHyLink.OnTapped = func() {
		fyne.CurrentApp().Clipboard().SetContent(appPath)
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
		widget.NewLabel("The Application is being installed at:"),
		container.NewHBox(
			appPathHyLink,
			copyHint,
		),
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
	)
}

func (a *ApplicationStep) OnInit() {
	a.Stepper.Footer.Hint = widget.NewLabel("Click Next to verify application.")
	err := goapp.CopyAppToConfigFolder()
	if err != nil {
		fmt.Println(err)
	}
}

func (a *ApplicationStep) OnNext() bool {
	// do nothing
	return true
}

func (a *ApplicationStep) OnPrevious() bool {
	// do nothing
	return true
}
