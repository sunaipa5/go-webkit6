package webkit

func WebViewGetSettings(webView uintptr) uintptr {
	return webkitWebViewGetSettings(webView)
}

func SettingsSetUserAgent(settings uintptr, userAgent string) {
	userAgentC := cstring(userAgent)
	webkitSettingsSetUserAgent(settings, userAgentC)
}

func SettingsSetEnableWebAudio(settings uintptr, enable bool) {
	webkitSettingsSetEnableWebAudio(settings, enable)
}

func SettingsSetEnableJavascript(settings uintptr, enable bool) {
	webkitSettingsSetEnableJavascript(settings, enable)
}

func SettingsSetJavascriptCanAccessClipboard(settings uintptr, enable bool) {
	webkitSettingsSetJavascriptCanAccessClipboard(settings, enable)
}

func SettingsSetEnableWebGL(settings uintptr, enable bool) {
	webkitSettingsSetEnableWebGL(settings, enable)
}

func SettingsSetEnablePageCache(settings uintptr, enable bool) {
	webkitSettingsSetEnablePageCache(settings, enable)
}

func SettingsSetEnableSmoothScrolling(settings uintptr, enable bool) {
	webkitSettingsSetEnableSmoothScrolling(settings, enable)
}

func SettingsSetAutoLoadImages(settings uintptr, enable bool) {
	webkitSettingsSetAutoLoadImages(settings, enable)
}

func SettingsSetEnableDeveloperExtras(settings uintptr, enable bool) {
	webkitSettingsSetEnableDeveloperExtras(settings, enable)
}

/*
## List of types
WEBKIT_HARDWARE_ACCELERATION_POLICY_ALWAYS = 0
WEBKIT_HARDWARE_ACCELERATION_POLICY_NEVER = 1
*/
func SettingsSetHardwareAccelerationPolicy(settings uintptr, policy int) {
	webkitSettingsSetHardwareAccelerationPolicy(settings, policy)
}

func SettingsSetAllowFileAccessFromFileUrls(settings uintptr, allow bool) {
	webkitSettingsSetAllowFileAccessFromFileUrls(settings, allow)
}

func SettingsSetAllowUniversalAccessFromFileUrls(settings uintptr, allow bool) {
	webkitSettingsSetAllowUniversalAccessFromFileUrls(settings, allow)
}

func SettingsSetEnable2dCanvasAcceleration(settings uintptr, enable bool) {
	webkitSettingsSetEnable2dCanvasAcceleration(settings, enable)
}

func SettingsSetEnableHyperlinkAuditing(settings uintptr, enable bool) {
	webkitSettingsSetEnableHyperlinkAuditing(settings, enable)
}

func SettingsSetEnableMediaStream(settings uintptr, enable bool) {
	webkitSettingsSetEnableMediaStream(settings, enable)
}

func SettingsSetEnableMediaCapabilities(settings uintptr, enable bool) {
	webkitSettingsSetEnableMediaCapabilities(settings, enable)
}

func SettingsSetEnableEncryptedMedia(settings uintptr, enable bool) {
	webkitSettingsSetEnableEncryptedMedia(settings, enable)
}

func SettingsSetEnableSiteSpecificQuirks(settings uintptr, enable bool) {
	webkitSettingsSetEnableSiteSpecificQuirks(settings, enable)
}

func SettingsSetEnableResizableTextAreas(settings uintptr, enable bool) {
	webkitSettingsSetEnableResizableTextAreas(settings, enable)
}

func SettingsSetEnableTabsToLinks(settings uintptr, enable bool) {
	webkitSettingsSetEnableTabsToLinks(settings, enable)
}

func SettingsSetEnableBackForwardNavigationGestures(settings uintptr, enable bool) {
	webkitSettingsSetEnableBackForwardNavigationGestures(settings, enable)
}

func SettingsSetEnableWriteConsoleMessagesToStdout(settings uintptr, enable bool) {
	webkitSettingsSetEnableWriteConsoleMessagesToStdout(settings, enable)
}
