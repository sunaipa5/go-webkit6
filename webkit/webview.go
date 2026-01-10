package webkit

import (
	"syscall"
)

func WebViewNew() uintptr {
	return webkit_web_view_new()
}

func WebViewLoadUri(webView uintptr, uri string) {
	cstr := append([]byte(uri), 0)
	webkit_web_view_load_uri(webView, &cstr[0])
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
	webkit_web_view_load_html(webView, c, b)
}

func WebsiteDataManagerGetType() uintptr {
	return webkit_website_data_manager_get_type()
}

func WebContextGetType() uintptr {
	return webkit_web_context_get_type()
}

func WebViewGetType() uintptr {
	return webkit_web_view_get_type()
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

	return webkit_security_origin_new(p, h, port)
}

func SecurityOriginNewForUri(uri string) uintptr {
	return webkit_security_origin_new_for_uri(cstring(uri))
}

// Network session
func NetworkSessionNew(dataDir, cacheDir string) uintptr {
	return webkit_network_session_new(&[]byte(dataDir)[0], &[]byte(cacheDir)[0])
}

func NetworkSessionGetDefault() uintptr {
	return webkit_network_session_get_default()
}

func NetworkSessionIsEphemeral(session uintptr) bool {
	return webkit_network_session_is_ephemeral(session)
}

func WebContextNew() uintptr {
	return webkit_web_context_new()
}

/*
context             uintptr WebKitWebContext

allowed_origins     uintptr GList

disallowed_origins  uintptr GList
*/

func WebContextInitializeNotificationPermissions(context uintptr, allowed_origins, disallowed_origins uintptr) {
	webkit_web_context_initialize_notification_permissions(
		context,
		allowed_origins,
		disallowed_origins,
	)
}
func WebViewTryClose(webView uintptr) {
	webkit_web_view_try_close(webView)
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

	webkit_web_view_evaluate_javascript(
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
	return webkit_navigation_policy_decision_get_navigation_action(decision)
}

func NavigationActionGetRequest(navAction uintptr) uintptr {
	return webkit_navigation_action_get_request(navAction)
}

func UriRequestGetUri(request uintptr) string {
	cstr := webkit_uri_request_get_uri(request)
	gostr := gostring(cstr)

	return gostr
}

//Content Manager

type ContentManager struct {
	Ptr uintptr
}

func NewUserContentManager() *ContentManager {
	managerPtr := webkit_user_content_manager_new()
	return &ContentManager{
		Ptr: managerPtr,
	}
}

func (m *ContentManager) UserContentManagerRegisterScriptMessageHandler(name string, world_name string) bool {
	status := webkit_user_content_manager_register_script_message_handler(m.Ptr, cstring(name), cstring(world_name))
	if status == 1 {
		return true
	}

	return false
}

func (m *ContentManager) UserContentManagerRegisterScriptMessageHandlerWithReply(name string, world_name string) bool {
	status := webkit_user_content_manager_register_script_message_handler_with_reply(m.Ptr, cstring(name), cstring(world_name))
	if status == 1 {
		return true
	}

	return false
}
