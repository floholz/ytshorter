package stepper

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/floholz/ytshorter/installer/internal/utils"
)

type MissingFilesPage struct {
	OnDismiss func()
}

func (m *MissingFilesPage) Show(extensionError, appError error) fyne.CanvasObject {
	header := widget.NewLabel("You are missing some Files!")
	header.Alignment = fyne.TextAlignCenter
	header.TextStyle.Bold = true

	extensionCheckMark := widget.NewCheck("good", nil)
	extFilesGood := extensionError == nil
	extensionCheckMark.Checked = extFilesGood
	extensionCheckMark.OnChanged = func(checked bool) {
		extensionCheckMark.SetChecked(extFilesGood)
	}
	if !extFilesGood {
		extensionCheckMark.SetText("missing")
	}

	appCheckMark := widget.NewCheck("good", nil)
	appFilesGood := appError == nil
	appCheckMark.Checked = appFilesGood
	appCheckMark.OnChanged = func(checked bool) {
		appCheckMark.SetChecked(appFilesGood)
	}
	if !appFilesGood {
		appCheckMark.SetText("missing")
	}

	return container.New(
		layout.NewCustomPaddedLayout(40, 40, 40, 40),
		container.NewBorder(
			container.NewVBox(
				header,
				widget.NewSeparator(),
			),
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewButton("Ignore", m.OnDismiss),
			),
			nil, nil,
			container.NewHBox(
				layout.NewSpacer(),
				container.NewVBox(
					layout.NewSpacer(),
					widget.NewLabel("Please download the full installer archive from:"),
					widget.NewHyperlink(
						"https://github.com/floholz/ytshorter/releases/latest",
						&url.URL{Scheme: "https", Host: "github.com", Path: "/floholz/ytshorter/releases/latest"},
					),
					layout.NewSpacer(),
					widget.NewForm(
						widget.NewFormItem("Browser Extension Files:", extensionCheckMark),
						widget.NewFormItem("Application Executable:", appCheckMark),
					),
					layout.NewSpacer(),
				),
				layout.NewSpacer(),
			),
		),
	)
}

func (m *MissingFilesPage) CheckAndShow(dismissContent *fyne.Container) fyne.CanvasObject {
	errExt := utils.CopyExtensionToConfigFolder()
	errApp := utils.CopyAppToConfigFolder()

	if errExt != nil || errApp != nil {
		fmt.Println("ext:", errExt)
		fmt.Println("app:", errApp)
		return m.Show(errExt, errApp)
	}

	return dismissContent
}
