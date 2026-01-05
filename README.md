# YTShorter (WIP)

YTShorter is a project designed to enhance the YouTube Shorts experience by allowing users to navigate between Shorts using global hotkeys. It consists of a Chrome extension, a Go-based native messaging host, and an installer.

## Project Structure

- `app/`: The Go native host application. It runs as a system tray app and listens for global hotkeys.
- `ext/`: The Chrome extension that communicates with the Go app and interacts with the YouTube page.
- `installer/`: A Fyne-based GUI installer to set up the Native Messaging manifest.
- `assets/`: Shared assets like icons.

## Architecture

```
Web Page (YouTube Shorts)
   ▲
   │ DOM access (click "Next")
Content Script (ext/content.js)
   ▲
   │ chrome.runtime.sendMessage
Background Service Worker (ext/background.js)
   ▲
   │ Native Messaging (stdin/stdout)
Go Native Host (app/main.go) [launched by Chrome]
   ▲
   │ Global Hotkey (robotgo/hook)
Keyboard Event (Ctrl+Shift+U)
```

## Setup & Installation (Development)

### 1. Build the Go App
```bash
cd app
go build -o ytshorter-app
```

### 2. Build the Installer
```bash
cd installer
go build -o installer-app
```

### 3. Load the Chrome Extension
1. Open Chrome and go to `chrome://extensions`.
2. Enable **Developer mode**.
3. Click **Load unpacked** and select the `ext/` folder in this repository.
4. Note the **Extension ID** (e.g., `abcdefghijklmnopqrstuvwxyzabcdef`).

### 4. Run the Installer
1. Run the `installer-app` built in step 2.
2. Paste the **Extension ID** into the input field.
3. Click **Install Native Host Manifest**.

### 5. Usage
1. Open a YouTube Shorts page (e.g., `https://www.youtube.com/shorts/...`).
2. The Go app should start automatically (indicated by a tray icon).
3. Press **Ctrl+Shift+U** (default) to skip to the next Short.
4. You can change the hotkey via the "Set Keybind" option in the tray menu.

## Technologies Used

- **Go**: For the native host and installer.
- **Fyne**: GUI toolkit for the tray app and installer.
- **RobotGo (hook)**: For global hotkey listening.
- **Chrome Native Messaging**: For communication between the extension and the Go app.
- **JavaScript (Manifest V3)**: For the Chrome extension.

## Current Status (WIP)

- [x] Basic architecture implemented.
- [x] Native messaging bridge working.
- [x] Global hotkey support.
- [x] System tray indicator.
- [x] Basic installer for manifest generation.
- [x] Support for YouTube SPA navigation (`yt-navigate-finish`).
- [ ] Windows support for the installer (Registry keys).
- [ ] Improved error handling and auto-reconnect.
- [ ] More "reactions" and actions from the tray app.

## License
GPLv3
