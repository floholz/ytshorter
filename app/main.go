package main

import (
	_ "embed"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	hook "github.com/robotn/gohook"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed assets/logo.png
var iconData []byte

var (
	activeKeybind = []string{"ctrl", "shift", "u"}
	hookStarted   = false
)

type Message struct {
	Type   string `json:"type"`
	Action string `json:"action,omitempty"`
}

func readMessage(r io.Reader) (*Message, error) {
	var length uint32
	if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
		return nil, err
	}

	payload := make([]byte, length)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}

	var msg Message
	if err := json.Unmarshal(payload, &msg); err != nil {
		return nil, err
	}

	return &msg, nil
}

func writeMessage(w io.Writer, msg Message) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, uint32(len(payload))); err != nil {
		return err
	}

	_, err = w.Write(payload)
	return err
}

func registerKeyHook(keybinding []string) {
	if hookStarted {
		hook.End()
	}

	// Log to stderr because stdout is for native messaging
	keybindingStr := strings.Join(keybinding, "+")
	keybindingStr = cases.Title(language.English).String(keybindingStr)
	_, _ = fmt.Fprintf(os.Stderr, "Registering hotkey: %s\n", keybindingStr)

	hook.Register(hook.KeyDown, keybinding, func(e hook.Event) {
		msg := Message{
			Type:   "HOTKEY_EVENT",
			Action: "next_short",
		}
		_ = writeMessage(os.Stdout, msg)
	})
	activeKeybind = keybinding

	s := hook.Start()
	hookStarted = true

	go func() {
		<-hook.Process(s)
	}()
}

func setKeybind(a fyne.App) {
	w := a.NewWindow("Set Keybind")
	w.Resize(fyne.NewSize(300, 100))
	w.SetFixedSize(true)
	icon := fyne.NewStaticResource("icon", iconData)
	w.SetIcon(icon)

	keybindingStr := strings.Join(activeKeybind, "+")
	keybindingStr = cases.Title(language.English).String(keybindingStr)
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
		registerKeyHook(newBind)
		w.Close()
	})

	w.CenterOnScreen()
	w.Show()
}

func main() {
	a := app.NewWithID("com.floholz.ytshorter")

	if desk, ok := a.(desktop.App); ok {
		headerItem := fyne.NewMenuItem("YTShorter is running", nil)
		headerItem.Disabled = true
		m := fyne.NewMenu("YTShorter",
			headerItem,
			fyne.NewMenuItem("Set Keybind", func() {
				setKeybind(a)
			}),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		)
		desk.SetSystemTrayMenu(m)
		icon := fyne.NewStaticResource("icon", iconData)
		desk.SetSystemTrayIcon(icon)
	}

	// Start Native Messaging Loop
	go func() {
		for {
			msg, err := readMessage(os.Stdin)
			if err != nil {
				if err == io.EOF {
					a.Quit()
					return
				}
				continue
			}

			if msg.Type == "EXTENSION_ACTION" {
				response := Message{
					Type:   "HOST_ACTION",
					Action: "reactions_comming_soon",
				}
				_ = writeMessage(os.Stdout, response)
			}
		}
	}()

	// Start Global Hotkey Listener
	registerKeyHook(activeKeybind)

	a.Run()
}
