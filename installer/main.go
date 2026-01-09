package main

import (
	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/floholz/ytshorter/installer/assets"
	"github.com/floholz/ytshorter/installer/internal/stepper"
)

func main() {
	a := app.NewWithID("com.floholz.ytshorter_installer")
	a.SetIcon(fyne.NewStaticResource("logo", assets.LogoData))
	w := a.NewWindow("YT Shorter Installer")

	installStepper := stepper.NewStepper()
	missingFiles := stepper.MissingFilesPage{OnDismiss: func() {
		w.SetContent(installStepper.Content)
	}}
	w.SetContent(missingFiles.CheckAndShow(installStepper.Content))

	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.ShowAndRun()
}
