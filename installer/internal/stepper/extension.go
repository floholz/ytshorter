package stepper

import (
	"fmt"
	"net/url"
	"path"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/floholz/ytshorter/installer/internal/chromiumext"
	"github.com/floholz/ytshorter/installer/internal/utils"
)

type ExtensionStep struct {
	Stepper *TStepper
	error   bool
}

func (e *ExtensionStep) Title() string {
	return "Install Browser Extension"
}

func (e *ExtensionStep) Content() fyne.CanvasObject {
	configPath, err := utils.GetConfigPath()
	if err != nil {
		configPath = "???"
	}
	extPath := path.Join(configPath, "browser-extension")

	copyHint := widget.NewLabelWithStyle("(click to copy)", fyne.TextAlignTrailing, fyne.TextStyle{Italic: true})

	extPathHyLink := widget.NewHyperlink(extPath, nil)
	extPathHyLink.OnTapped = func() {
		fyne.CurrentApp().Clipboard().SetContent(extPath)
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

	image := canvas.NewImageFromFile(path.Join("assets", "screenshot-browser_extension-anotated.png"))
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(300, 150))

	return container.NewVBox(
		container.NewHBox(
			container.NewVBox(
				layout.NewSpacer(),
				widget.NewForm(
					widget.NewFormItem("1.",
						container.NewHBox(
							widget.NewLabel("Open "),
							widget.NewHyperlink("chrome://extensions", &url.URL{Scheme: "chrome", Host: "extensions"}),
						),
					),
					widget.NewFormItem("2.", widget.NewLabel("Enable 'Developer mode'")),
					widget.NewFormItem("3.", widget.NewLabel("Click 'Load unpacked'")),
				),
			),
			layout.NewSpacer(),
			image,
		),
		widget.NewForm(
			widget.NewFormItem("4.",
				container.NewVBox(
					widget.NewLabel("Select the 'browser-extension' folder from"),
					container.NewHBox(
						extPathHyLink,
						copyHint,
					),
				),
			),
		),
	)
}

func (e *ExtensionStep) OnInit() {
	e.Stepper.Footer.Hint = widget.NewLabel("Click Next to verify installation.")
	err := chromiumext.CopyExtensionToConfigFolder()
	if err != nil {
		fmt.Println(err)
	}
}

func (e *ExtensionStep) OnNext() bool {
	if e.error {
		return true // continue after second next click
	}

	ext, err := chromiumext.Detect("mghmjdfcifpdodkfdggjelopdfopgale", "Default")
	if err != nil || ext == nil || len(ext.DisableReasons) != 0 {
		e.Stepper.Footer.Hint = widget.NewLabelWithStyle("Installation of the chrome extension could not be verified!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: true})
		e.Stepper.Footer.Hint.Show()
		e.error = true
		fmt.Println(err, ext)
		return false
	}

	return true
}

func (e *ExtensionStep) OnPrevious() bool {
	// do nothing
	return true
}
