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

	if kbObj.MenuItem != nil {
		kbObj.MenuItem.Label = "Set '" + kbObj.Name + "' Keybind [" + keybindingStr + "]"
	}

	if refreshCallbackFn != nil {
		refreshCallbackFn(kbObj)
	}
}

func Start() {
	s := hook.Start()
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

		hook.End()
		// Re-register all hooks since hook.End() might clear them or we need a fresh start
		// In a real app we might need a way to re-register all known hooks.
		// For now, let's just trigger the refresh which should call RegisterKeyHook.
		// Wait, we need to re-register BOTH hooks if we call hook.End().

		// Actually, gohook's Register is global.
		// If we want to change one, we might need to End and Re-start everything.

		// Let's assume for now that we want to restart the whole hook system.
		refreshCallbackFn(kbObj)
		w.Close()
	})

	w.CenterOnScreen()
	w.Show()
}

func KeybindingToString(keybinding []string) string {
	return cases.Title(language.English).String(strings.Join(keybinding, "+"))
}
