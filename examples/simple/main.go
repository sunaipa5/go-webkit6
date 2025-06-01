package main

import (
	"fmt"
	"os"
	"path"
	"unsafe"

	"github.com/sunaipa5/go-webkit6/webkit"

	"github.com/jwijenbergh/puregotk/v4/adw"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
	"github.com/jwijenbergh/puregotk/v4/gtk"
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
	userHome, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	baseDataDir := path.Join(userHome, "webkit6-test", "data")
	baseCacheDir := path.Join(userHome, "webkit6-test", "cache")

	/*
		Create a new network session with base data dir and base cache dir
	*/
	netsession := webkit.NetworkSessionNew(baseDataDir, baseCacheDir)

	/*
	   Create a new GObject instance of type WebView.
	   Use `WebViewGetType` to obtain the type ID.

	   Assign the network session to the `network-session` property.
	   You do not need to convert the network session to a GObject manually,
	   as it is already a GObject-compatible pointer.
	*/
	webviewObj := gobject.NewObject(
		gobject.Type(webkit.WebViewGetType()),
		"network-session", netsession,
	)

	webview := webviewObj.Ptr

	webviewWidget := gtk.WidgetNewFromInternalPtr(webview)

	webkitsettings := webkit.WebViewGetSettings(webview)
	webkit.SettingsSetEnablePageCache(webkitsettings, true)
	webkit.SettingsSetEnableWebGL(webkitsettings, true)
	webkit.SettingsSetEnableSmoothScrolling(webkitsettings, true)
	webkit.SettingsSetEnableWebAudio(webkitsettings, true)
	webkit.SettingsSetJavascriptCanAccessClipboard(webkitsettings, false)
	webkit.SettingsSetHardwareAccelerationPolicy(webkitsettings, 0)

	//DEVTOOLS
	webkit.SettingsSetEnableDeveloperExtras(webkitsettings, true)

	webkit.WebViewLoadUri(webview, "https://go.dev")

	//Permission request handler
	permissionFunc := func(webview, request, userData uintptr) {
		fmt.Println("Permission request")

		//Handle notification request
		if gobject.TypeCheckInstanceIsA((*gobject.TypeInstance)(unsafe.Pointer(request)), types.GType(webkit.NotificationPermissionRequestGetType())) {
			//Allow permission the notification request.
			//After this, the notification will be shown.
			webkit.PermissionRequestAllow(request)
		}
	}

	//Create signal connection
	gobject.SignalConnect(webview, "permission-request", glib.NewCallback(&permissionFunc))

	return webviewWidget
}
