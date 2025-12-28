package webkit

import (
	"os"

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

	webkitSettingsSetAllowFileAccessFromFileUrls         func(uintptr, bool)
	webkitSettingsSetAllowUniversalAccessFromFileUrls    func(uintptr, bool)
	webkitSettingsSetEnable2dCanvasAcceleration          func(uintptr, bool)
	webkitSettingsSetEnableHyperlinkAuditing             func(uintptr, bool)
	webkitSettingsSetEnableMediaStream                   func(uintptr, bool)
	webkitSettingsSetEnableMediaCapabilities             func(uintptr, bool)
	webkitSettingsSetEnableEncryptedMedia                func(uintptr, bool)
	webkitSettingsSetEnableSiteSpecificQuirks            func(uintptr, bool)
	webkitSettingsSetEnableResizableTextAreas            func(uintptr, bool)
	webkitSettingsSetEnableTabsToLinks                   func(uintptr, bool)
	webkitSettingsSetEnableBackForwardNavigationGestures func(uintptr, bool)
	webkitSettingsSetEnableWriteConsoleMessagesToStdout  func(uintptr, bool)

	//permission
	webkit_notification_permission_request_get_type func() uintptr
	webkit_geolocation_permission_request_get_type  func() uintptr
	webkit_permission_request_allow                 func(uintptr)

	//Content Manager
	webkit_user_content_manager_new                                        func() uintptr
	webkit_user_content_manager_register_script_message_handler            func(manager uintptr, name uintptr, world uintptr) int32
	webkit_user_content_manager_register_script_message_handler_with_reply func(manager uintptr, name uintptr, world uintptr) int32
)

// You can change the library location with ldflag or set the WEBKITGTK_PATH environment variable.
// -X main.WebKitGTKPath=/custom/path/libwebkitgtk-6.0.so.4
// Don't change this variable while the application is running
var WebKitGTKPath string

func init() {
	if WebKitGTKPath == "" {
		if env := os.Getenv("WEBKITGTK_PATH"); env != "" {
			WebKitGTKPath = env
		} else {
			WebKitGTKPath = "libwebkitgtk-6.0.so.4"
		}
	}

	lib, err := purego.Dlopen(WebKitGTKPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
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

	purego.RegisterLibFunc(&webkitSettingsSetAllowFileAccessFromFileUrls, lib, "webkit_settings_set_allow_file_access_from_file_urls")
	purego.RegisterLibFunc(&webkitSettingsSetAllowUniversalAccessFromFileUrls, lib, "webkit_settings_set_allow_universal_access_from_file_urls")
	purego.RegisterLibFunc(&webkitSettingsSetEnable2dCanvasAcceleration, lib, "webkit_settings_set_enable_2d_canvas_acceleration")
	purego.RegisterLibFunc(&webkitSettingsSetEnableHyperlinkAuditing, lib, "webkit_settings_set_enable_hyperlink_auditing")
	purego.RegisterLibFunc(&webkitSettingsSetEnableMediaStream, lib, "webkit_settings_set_enable_media_stream")
	purego.RegisterLibFunc(&webkitSettingsSetEnableMediaCapabilities, lib, "webkit_settings_set_enable_media_capabilities")
	purego.RegisterLibFunc(&webkitSettingsSetEnableEncryptedMedia, lib, "webkit_settings_set_enable_encrypted_media")
	purego.RegisterLibFunc(&webkitSettingsSetEnableSiteSpecificQuirks, lib, "webkit_settings_set_enable_site_specific_quirks")
	purego.RegisterLibFunc(&webkitSettingsSetEnableResizableTextAreas, lib, "webkit_settings_set_enable_resizable_text_areas")
	purego.RegisterLibFunc(&webkitSettingsSetEnableTabsToLinks, lib, "webkit_settings_set_enable_tabs_to_links")
	purego.RegisterLibFunc(&webkitSettingsSetEnableBackForwardNavigationGestures, lib, "webkit_settings_set_enable_back_forward_navigation_gestures")
	purego.RegisterLibFunc(&webkitSettingsSetEnableWriteConsoleMessagesToStdout, lib, "webkit_settings_set_enable_write_console_messages_to_stdout")

	//permission
	purego.RegisterLibFunc(&webkit_permission_request_allow, lib, "webkit_permission_request_allow")
	purego.RegisterLibFunc(&webkit_notification_permission_request_get_type, lib, "webkit_notification_permission_request_get_type")
	purego.RegisterLibFunc(&webkit_geolocation_permission_request_get_type, lib, "webkit_geolocation_permission_request_get_type")

	//Content manager
	purego.RegisterLibFunc(&webkit_user_content_manager_new, lib, "webkit_user_content_manager_new")
	purego.RegisterLibFunc(&webkit_user_content_manager_register_script_message_handler, lib, "webkit_user_content_manager_register_script_message_handler")
	purego.RegisterLibFunc(&webkit_user_content_manager_register_script_message_handler_with_reply, lib, "webkit_user_content_manager_register_script_message_handler_with_reply")

}
