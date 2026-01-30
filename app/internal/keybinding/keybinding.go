package keybinding

import (
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/floholz/ytshorter/app/internal/messaging"
	hook "github.com/robotn/gohook"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type KeyBindObject struct {
	Name        string
	Keybind     []string
	Action      string
	MenuItem    *fyne.MenuItem
	hookStarted bool
}

func RegisterKeyHook(kbObj *KeyBindObject, refreshCallbackFn func(kbObj *KeyBindObject)) {
	if kbObj.hookStarted {
		hook.End()
	}

	// Log to stderr because stdout is for native messaging
	keybindingStr := KeybindingToString(kbObj.Keybind)
	_, _ = fmt.Fprintf(os.Stderr, "Registering hotkey: %s\n", keybindingStr)

	hook.Register(hook.KeyDown, kbObj.Keybind, func(e hook.Event) {
		msg := messaging.Message{
			Type:   "HOTKEY_EVENT",
			Action: kbObj.Action,
		}
		_ = messaging.WriteMessage(os.Stdout, msg)
	})

	kbObj.MenuItem.Label = "Set '" + kbObj.Name + "' Keybind [" + keybindingStr + "]"
	refreshCallbackFn(kbObj)

	s := hook.Start()
	kbObj.hookStarted = true

	go func() {
		<-hook.Process(s)
	}()
}

func SetKeybind(a fyne.App, kbObj *KeyBindObject, refreshCallbackFn func(kbObj *KeyBindObject)) {
	w := a.NewWindow("Set '" + kbObj.Name + "' Keybind")
	w.Resize(fyne.NewSize(300, 100))
	w.SetFixedSize(true)
	w.SetIcon(theme.SettingsIcon())

	keybindingStr := KeybindingToString(kbObj.Keybind)
	currentLabel := widget.NewLabel("Current keybind: (" + keybindingStr + ")")
	currentLabel.Alignment = fyne.TextAlignCenter

	label := widget.NewLabel("Press your new key (Ctrl+Shift+...)")
	label.Alignment = fyne.TextAlignCenter

	w.SetContent(container.NewVBox(
		currentLabel,
		label,
		widget.NewButton("Cancel", func() {
			w.Close()
		}),
	))

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		newKey := strings.ToLower(string(k.Name))
		// Fyne key names might need mapping to robotgo hook names if they differ
		// For simple letters it should be fine.

		newBind := []string{"ctrl", "shift", newKey}
		kbObj.Keybind = newBind
		RegisterKeyHook(kbObj, refreshCallbackFn)
		w.Close()
	})

	w.CenterOnScreen()
	w.Show()
}

func KeybindingToString(keybinding []string) string {
	return cases.Title(language.English).String(strings.Join(keybinding, "+"))
}
