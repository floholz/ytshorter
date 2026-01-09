package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/floholz/ytshorter/installer/internal/stepper"
)

//go:embed assets/logo.png
var iconData []byte

type NativeManifest struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Path           string   `json:"path"`
	Type           string   `json:"type"`
	AllowedOrigins []string `json:"allowed_origins"`
}

func main() {
	a := app.NewWithID("com.floholz.ytshorter_installer")
	a.SetIcon(fyne.NewStaticResource("logo", iconData))
	w := a.NewWindow("YT Shorter Installer")

	w.SetContent(stepper.NewStepper().Content)

	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.ShowAndRun()
}

func _main() {
	myApp := app.New()
	myApp.SetIcon(fyne.NewStaticResource("logo", iconData))
	myWindow := myApp.NewWindow("YT Shorter Installer")

	instruction := widget.NewLabel("1. Open chrome://extensions\n2. Enable Developer Mode\n3. Click 'Load unpacked' and select the 'ext' folder\n4. Copy the Extension ID and paste it below:")

	extensionIDEntry := widget.NewEntry()
	extensionIDEntry.SetPlaceHolder("Extension ID")

	statusLabel := widget.NewLabel("")

	installBtn := widget.NewButton("Install Native Host Manifest", func() {
		extID := extensionIDEntry.Text
		if extID == "" {
			statusLabel.SetText("Please enter the Extension ID")
			return
		}

		err := installManifest(extID)
		if err != nil {
			statusLabel.SetText(fmt.Sprintf("Error: %v", err))
		} else {
			statusLabel.SetText("Manifest installed successfully!")
		}
	})

	openChromeBtn := widget.NewButton("Open Chrome Extensions", func() {
		openBrowser("chrome://extensions")
	})

	myWindow.SetContent(container.NewVBox(
		instruction,
		openChromeBtn,
		extensionIDEntry,
		installBtn,
		statusLabel,
	))

	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}

func installManifest(extID string) error {
	hostName := "com.floholz.ytshorter"

	// Determine app path - assuming it's built and placed in the app/ directory
	// In a real scenario, we might want to move it to a standard location
	absAppPath, err := filepath.Abs("../app/ytshorter-app")
	if err != nil {
		return err
	}

	manifest := NativeManifest{
		Name:        hostName,
		Description: "YT Shorter Go native host",
		Path:        absAppPath,
		Type:        "stdio",
		AllowedOrigins: []string{
			fmt.Sprintf("chrome-extension://%s/", extID),
		},
	}

	manifestBytes, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}

	var manifestPath string
	home, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "linux":
		manifestPath = filepath.Join(home, ".config/google-chrome/NativeMessagingHosts", hostName+".json")
	case "darwin":
		manifestPath = filepath.Join(home, "Library/Application Support/Google/Chrome/NativeMessagingHosts", hostName+".json")
	case "windows":
		// Windows requires registry keys, which is more complex.
		// For now, let's focus on Linux/Darwin or provide a warning.
		return fmt.Errorf("windows support not implemented in this demo")
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	err = os.MkdirAll(filepath.Dir(manifestPath), 0755)
	if err != nil {
		return err
	}

	return os.WriteFile(manifestPath, manifestBytes, 0644)
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		fmt.Printf("Failed to open browser: %v\n", err)
	}
}
