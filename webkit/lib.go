package webkit

import (
	"github.com/jwijenbergh/purego"
)

var (
	webkitWebViewNew                func() uintptr
	webkitWebViewGetType            func() uintptr
	webkitWebViewLoadURI            func(uintptr, *byte)
	webkitWebsiteDataManagerGetType func() uintptr
	webkitWebViewTryClose           func(uintptr)
	webkitWebViewLoadHtml           func(uintptr, content uintptr, base_uri uintptr)
	webkitWebViewEvaluateJavascript func(
		webView uintptr,
		script uintptr,
		length int64,
		cancellable uintptr,
		callback uintptr,
		userData uintptr,
		userDataDestroy uintptr,
		callbackData uintptr,
	)

	//Security
	webkitSecurityOriginNew       func(protocol, host *byte, port uint16) uintptr
	webkitSecurityOriginNewForUri func(uintptr) uintptr

	//Navigation
	webkitNavigationPolicyDecisionGetNavigationAction func(uintptr) uintptr
	webkitNavigationActionGetRequest                  func(uintptr) uintptr
	webkitUriRequestGetUri                            func(uintptr) uintptr

	//Network
	webkitNetworkSessionNew         func(dataDir, cacheDir *byte) uintptr
	webkitNetworkSessionGetDefault  func() uintptr
	webkitNetworkSessionIsEphemeral func(uintptr) bool

	//Context
	webkitWebContextNew                               func() uintptr
	webkitWebContextGetType                           func() uintptr
	webkitWebContextInitializeNotificationPermissions func(
		context uintptr,
		allowedOrigins uintptr,
		disallowedOrigins uintptr,
	)

	//Settings
	webkitWebViewGetSettings                      func(webview uintptr) uintptr
	webkitSettingsSetUserAgent                    func(settings uintptr, userAgent uintptr)
	webkitSettingsSetEnableWebAudio               func(uintptr, bool)
	webkitSettingsSetEnableJavascript             func(uintptr, bool)
	webkitSettingsSetJavascriptCanAccessClipboard func(uintptr, bool)
	webkitSettingsSetEnableWebGL                  func(uintptr, bool)
	webkitSettingsSetEnablePageCache              func(uintptr, bool)
	webkitSettingsSetEnableSmoothScrolling        func(uintptr, bool)
	webkitSettingsSetAutoLoadImages               func(uintptr, bool)
	webkitSettingsSetHardwareAccelerationPolicy   func(uintptr, int)
	webkitSettingsSetEnableDeveloperExtras        func(uintptr, bool)

	//permission
	webkit_notification_permission_request_get_type func() uintptr
	webkit_geolocation_permission_request_get_type  func() uintptr
	webkit_permission_request_allow                 func(uintptr)
)

func init() {
	lib, err := purego.Dlopen("libwebkitgtk-6.0.so", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&webkitWebViewNew, lib, "webkit_web_view_new")
	purego.RegisterLibFunc(&webkitWebViewGetType, lib, "webkit_web_view_get_type")
	purego.RegisterLibFunc(&webkitWebViewLoadURI, lib, "webkit_web_view_load_uri")
	purego.RegisterLibFunc(&webkitWebsiteDataManagerGetType, lib, "webkit_website_data_manager_get_type")
	purego.RegisterLibFunc(&webkitWebViewTryClose, lib, "webkit_web_view_try_close")
	purego.RegisterLibFunc(&webkitWebViewLoadHtml, lib, "webkit_web_view_load_html")
	purego.RegisterLibFunc(&webkitWebViewEvaluateJavascript, lib, "webkit_web_view_evaluate_javascript")

	//Security
	purego.RegisterLibFunc(&webkitSecurityOriginNew, lib, "webkit_security_origin_new")
	purego.RegisterLibFunc(&webkitSecurityOriginNewForUri, lib, "webkit_security_origin_new_for_uri")

	//Navigation
	purego.RegisterLibFunc(&webkitNavigationPolicyDecisionGetNavigationAction, lib, "webkit_navigation_policy_decision_get_navigation_action")
	purego.RegisterLibFunc(&webkitNavigationActionGetRequest, lib, "webkit_navigation_action_get_request")
	purego.RegisterLibFunc(&webkitUriRequestGetUri, lib, "webkit_uri_request_get_uri")

	//Context
	purego.RegisterLibFunc(&webkitWebContextNew, lib, "webkit_web_context_new")
	purego.RegisterLibFunc(&webkitWebContextGetType, lib, "webkit_web_context_get_type")

	//Network Session
	purego.RegisterLibFunc(&webkitNetworkSessionNew, lib, "webkit_network_session_new")
	purego.RegisterLibFunc(&webkitNetworkSessionGetDefault, lib, "webkit_network_session_get_default")
	purego.RegisterLibFunc(&webkitNetworkSessionIsEphemeral, lib, "webkit_network_session_is_ephemeral")
	purego.RegisterLibFunc(&webkitWebContextInitializeNotificationPermissions, lib, "webkit_web_context_initialize_notification_permissions")

	//Settings
	purego.RegisterLibFunc(&webkitWebViewGetSettings, lib, "webkit_web_view_get_settings")
	purego.RegisterLibFunc(&webkitSettingsSetUserAgent, lib, "webkit_settings_set_user_agent")
	purego.RegisterLibFunc(&webkitSettingsSetEnableWebAudio, lib, "webkit_settings_set_enable_webaudio")
	purego.RegisterLibFunc(&webkitSettingsSetEnableJavascript, lib, "webkit_settings_set_enable_javascript")
	purego.RegisterLibFunc(&webkitSettingsSetJavascriptCanAccessClipboard, lib, "webkit_settings_set_javascript_can_access_clipboard")
	purego.RegisterLibFunc(&webkitSettingsSetEnableWebGL, lib, "webkit_settings_set_enable_webgl")
	purego.RegisterLibFunc(&webkitSettingsSetEnablePageCache, lib, "webkit_settings_set_enable_page_cache")
	purego.RegisterLibFunc(&webkitSettingsSetEnableSmoothScrolling, lib, "webkit_settings_set_enable_smooth_scrolling")
	purego.RegisterLibFunc(&webkitSettingsSetAutoLoadImages, lib, "webkit_settings_set_auto_load_images")
	purego.RegisterLibFunc(&webkitSettingsSetHardwareAccelerationPolicy, lib, "webkit_settings_set_hardware_acceleration_policy")
	purego.RegisterLibFunc(&webkitSettingsSetEnableDeveloperExtras, lib, "webkit_settings_set_enable_developer_extras")

	//permission
	purego.RegisterLibFunc(&webkit_permission_request_allow, lib, "webkit_permission_request_allow")
	purego.RegisterLibFunc(&webkit_notification_permission_request_get_type, lib, "webkit_notification_permission_request_get_type")
	purego.RegisterLibFunc(&webkit_geolocation_permission_request_get_type, lib, "webkit_geolocation_permission_request_get_type")

}
