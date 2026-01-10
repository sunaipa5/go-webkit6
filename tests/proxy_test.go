package tests

import (
	"path"
	"testing"

	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
	"github.com/sunaipa5/go-webkit6/webkit"
)

func TestProxy(t *testing.T) {
	baseTestApp(func() *gtk.Widget {
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
	})
}
