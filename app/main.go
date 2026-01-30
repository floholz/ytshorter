package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"github.com/floholz/ytshorter/app/internal"
	"github.com/floholz/ytshorter/app/internal/keybinding"
	"github.com/floholz/ytshorter/app/internal/messaging"
)

//go:embed assets/logo.png
var AppIconData []byte

var (
	kboNext     *keybinding.KeyBindObject
	kboPause    *keybinding.KeyBindObject
	SystrayMenu *fyne.Menu
	AppConfig   internal.Config
)

func main() {
	AppConfig = internal.LoadConfig()

	kboNext = &keybinding.KeyBindObject{
		Name:     "Next",
		Keybind:  AppConfig.Keybinds.Next,
		Action:   "next_short",
		MenuItem: nil,
	}
	kboPause = &keybinding.KeyBindObject{
		Name:     "Pause",
		Keybind:  AppConfig.Keybinds.Pause,
		Action:   "pause_short",
		MenuItem: nil,
	}

	a := app.NewWithID("com.floholz.ytshorter")

	if desk, ok := a.(desktop.App); ok {
		headerItem := fyne.NewMenuItem("YTShorter is active", nil)
		headerItem.Disabled = true

		kboNext.MenuItem = fyne.NewMenuItemWithIcon(
			"Set 'Next' Keybind ["+keybinding.KeybindingToString(AppConfig.Keybinds.Next)+"]",
			theme.ContentAddIcon(),
			func() {
				keybinding.SetKeybind(a, kboNext, saveConfigAndRefreshSystray)
			})
		kboPause.MenuItem = fyne.NewMenuItemWithIcon(
			"Set 'Pause' Keybind ["+keybinding.KeybindingToString(AppConfig.Keybinds.Pause)+"]",
			theme.ContentAddIcon(),
			func() {
				keybinding.SetKeybind(a, kboPause, saveConfigAndRefreshSystray)
			})

		quitItem := fyne.NewMenuItemWithIcon("Quit", theme.LogoutIcon(), func() {
			a.Quit()
		})
		quitItem.IsQuit = true

		SystrayMenu = fyne.NewMenu("YTShorter",
			headerItem,
			kboNext.MenuItem,
			kboPause.MenuItem,
			fyne.NewMenuItemSeparator(),
			quitItem,
		)
		desk.SetSystemTrayMenu(SystrayMenu)
		icon := fyne.NewStaticResource("icon", AppIconData)
		desk.SetSystemTrayIcon(icon)
	}

	// Start Native Messaging Loop
	go func() {
		for {
			msg, err := messaging.ReadMessage(os.Stdin)
			if err != nil {
				if err == io.EOF {
					a.Quit()
					return
				}
				continue
			}

			if msg.Type == "EXTENSION_ACTION" {
				response := messaging.Message{
					Type:   "HOST_ACTION",
					Action: "reactions_comming_soon",
				}
				_ = messaging.WriteMessage(os.Stdout, response)
			}
		}
	}()

	// Start Global Hotkey Listener
	keybinding.RegisterKeyHook(kboNext, saveConfigAndRefreshSystray)
	keybinding.RegisterKeyHook(kboPause, saveConfigAndRefreshSystray)

	a.Run()
}

func saveConfigAndRefreshSystray(kbo *keybinding.KeyBindObject) {
	AppConfig.SetKeybind(kbo.Keybind, kbo.Name)
	err := AppConfig.Save()
	if err != nil {
		fmt.Printf("Failed to save config: %v\n", err)
	}
	SystrayMenu.Refresh()
}
