package webkit

import (
	"syscall"
	"unsafe"
)

func WebViewNew() uintptr {
	return webkitWebViewNew()
}

func WebViewLoadUri(webView uintptr, uri string) {
	cstr := append([]byte(uri), 0)
	webkitWebViewLoadURI(webView, &cstr[0])
}

/*
baseUri is used for host

# Example

<a href="baseUri/info">link</a>

<img src="baseUri/car.png">

# Usage

WebViewLoadHtml(webview, `

	<!DOCTYPE html>

<html>
<head>

	<title>Hello WebkitGTK</title>

</head>
<body>

	<h1>Hello WebkitGTK</h1>

</body>
</html>

	`, "")
*/
func WebViewLoadHtml(webView uintptr, content, baseUri string) {
	c := cstring(content)
	b := cstring(baseUri)
	webkitWebViewLoadHtml(webView, c, b)
}

func WebsiteDataManagerGetType() uintptr {
	return webkitWebsiteDataManagerGetType()
}

func WebContextGetType() uintptr {
	return webkitWebContextGetType()
}

func WebViewGetType() uintptr {
	return webkitWebViewGetType()
}

// Security
func SecurityOriginNew(protocol, host string, port uint16) uintptr {
	cs := func(s string) *byte {
		ptr, err := syscall.BytePtrFromString(s)
		if err != nil {
			panic("[WEBKIT - SecurityOriginNew] " + "Failed to convert go string to cstring ERR:" + err.Error())
		}
		return ptr
	}

	p := cs(protocol)
	h := cs(protocol)

	return webkitSecurityOriginNew(p, h, port)
}

func SecurityOriginNewForUri(uri string) uintptr {
	return webkitSecurityOriginNewForUri(cstring(uri))
}

// Network session
func NetworkSessionNew(dataDir, cacheDir string) uintptr {
	return webkitNetworkSessionNew(&[]byte(dataDir)[0], &[]byte(cacheDir)[0])
}

func NetworkSessionGetDefault() uintptr {
	return webkitNetworkSessionGetDefault()
}

func NetworkSessionIsEphemeral(session uintptr) bool {
	return webkitNetworkSessionIsEphemeral(session)
}

func WebContextNew() uintptr {
	return webkitWebContextNew()
}

/*
context             uintptr WebKitWebContext

allowed_origins     uintptr GList

disallowed_origins  uintptr GList
*/

func WebContextInitializeNotificationPermissions(context uintptr, allowed_origins, disallowed_origins uintptr) {
	webkitWebContextInitializeNotificationPermissions(
		context,
		allowed_origins,
		disallowed_origins,
	)
}
func WebViewTryClose(webView uintptr) {
	webkitWebViewTryClose(webView)
}

/*
# Example

WebViewEvaluateJavascript(webview, `alert("JS test");`, 0, 0, 0, 0, 0, 0)

# Required

webView         WebKitWebView*

js              JavaScript code

# Optional arguments (pass 0 if unused):

cancellable     GCancellable*

callback        GAsyncReadyCallback

userData        User data pointer

userDataDestroy GDestroyNotify for userData

callbackData    Additional callback data
*/
func WebViewEvaluateJavascript(webView uintptr, js string, length int64, cancellable uintptr, callback uintptr, userData uintptr, userDataDestroy uintptr, callbackData uintptr) {
	jsPtr := cstring(js)

	webkitWebViewEvaluateJavascript(
		webView,
		jsPtr,
		int64(len(js)),
		cancellable,
		callback,
		userData,
		userDataDestroy,
		callbackData,
	)
}

func NavigationPolicyDecisionGetNavigationAction(decision uintptr) uintptr {
	return webkitNavigationPolicyDecisionGetNavigationAction(decision)
}

func NavigationActionGetRequest(navAction uintptr) uintptr {
	return webkitNavigationActionGetRequest(navAction)
}

func UriRequestGetUri(request uintptr) string {
	cstr := webkitUriRequestGetUri(request)
	gostr := gostring(unsafe.Pointer(cstr))

	return gostr
}
