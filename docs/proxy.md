# Proxy usage

```go
  baseDataDir := path.Join("./", "webkitgtk-test", "data")
  baseCacheDir := path.Join("./", "webkitgtk-test", "cache")

  //Create network session
	netsession := webkit.NetworkSessionNew(baseDataDir, baseCacheDir)


  //To ignore hosts
  //proxySettings := webkit.NetworkProxySettingsNew("socks://127.0.0.1:9050", []string{"check.torproject.org"})

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
```
