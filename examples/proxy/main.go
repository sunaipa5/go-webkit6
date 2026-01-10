package main

import (
	"os"
	"path"

	"github.com/jwijenbergh/puregotk/v4/adw"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
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

	baseDataDir := path.Join("./", "webkitgtk-test", "data")
	baseCacheDir := path.Join("./", "webkitgtk-test", "cache")

	netsession := webkit.NetworkSessionNew(baseDataDir, baseCacheDir)

	proxySettings := webkit.NetworkProxySettingsNew("socks://127.0.0.1:9050", nil)
	webkit.NetworkSessionSetProxySettings(netsession, webkit.WEBKIT_NETWORK_PROXY_MODE_CUSTOM, proxySettings)

	webviewObj := gobject.NewObject(
		gobject.Type(webkit.WebViewGetType()),
		"network-session", netsession,
	)

	webview := webviewObj.Ptr

	webViewWidget := gtk.WidgetNewFromInternalPtr(webview)

	settings := webkit.WebViewGetSettings(webviewObj.Ptr)
	webkit.SettingsSetEnableDeveloperExtras(settings, true)

	webkit.WebViewLoadUri(webview, "https://check.torproject.org/")

	return webViewWidget
}
