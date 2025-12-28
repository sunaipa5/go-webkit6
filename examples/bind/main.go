package main

import (
	"fmt"
	"os"

	"github.com/jwijenbergh/puregotk/v4/adw"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
	jsc "github.com/sunaipa5/go-webkit6/javascriptcore"
	"github.com/sunaipa5/go-webkit6/webkit"
)

func main() {
	adw.Init()

	app := adw.NewApplication("com.test.app", gio.GApplicationFlagsNoneValue)
	defer app.Unref()

	actcb := func(_ gio.Application) {
		activate(app)
	}
	app.ConnectActivate(&actcb)

	if code := app.Run(len(os.Args), os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *adw.Application) {
	window := adw.NewApplicationWindow(&app.Application)

	mainBox := gtk.NewBox(gtk.OrientationVerticalValue, 0)

	headerBar := adw.NewHeaderBar()
	headerBar.AddCssClass("flat")
	headerBar.AddCssClass("compact")
	headerBar.SetShowTitle(false)
	headerBar.SetShowEndTitleButtons(true)
	mainBox.Append(&headerBar.Widget)

	webViewWidget := init_webview()

	webViewWidget.SetVexpand(true)
	webViewWidget.SetHexpand(true)
	webViewWidget.Show()

	mainBox.Append(webViewWidget)

	window.SetContent(&mainBox.Widget)
	window.SetDefaultSize(800, 600)
	window.Present()

}

func init_webview() *gtk.Widget {

	manager := webkit.NewUserContentManager()

	manager.UserContentManagerRegisterScriptMessageHandler("greet", "")

	greet := func(
		self uintptr,
		JsResult uintptr,
		userData uintptr,
	) int32 {
		if jsc.ValueIsString(JsResult) {
			fmt.Println(jsc.ValueToString(JsResult))
		} else {
			fmt.Println("Value isn't string")
		}
		return 1
	}

	gobject.SignalConnect(
		manager.Ptr,
		"script-message-received::greet",
		glib.NewCallback(&greet),
	)

	webviewObj := gobject.NewObject(
		gobject.Type(webkit.WebViewGetType()),
		"user-content-manager", manager.Ptr,
	)

	webview := webviewObj.Ptr

	webviewWidget := gtk.WidgetNewFromInternalPtr(webview)

	settings := webkit.WebViewGetSettings(webviewObj.Ptr)
	webkit.SettingsSetEnableDeveloperExtras(settings, true)

	webkit.WebViewLoadHtml(webview, `
		<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>WebKitGTK JS Bind Test</title>
</head>
<body>
<h1>WebKitGTK6 JS Bridge Test</h1>
<button onclick="sendMessage()">Send Message to Go</button>
<script>
function sendMessage() {
console.log("send message triggered");
    window.webkit?.messageHandlers?.greet?.postMessage("Hello from JS!") ||
    window.openInEditor?.("Hello from JS!");
    console.log("send message done");
}
</script>
</body>
</html>
`, "app://main")

	return webviewWidget
}
