package assets

import _ "embed"

//go:embed logo.png
var LogoData []byte

//go:embed screenshot-browser_extension-annotated.png
var ScreenshotBrowserExtensionData []byte

var NativeHostName = "com.floholz.ytshorter"
var ExtensionId = "mghmjdfcifpdodkfdggjelopdfopgale"
