package webkit

import (
	"os"
	"unsafe"

	"github.com/jwijenbergh/purego"
)

var (
	webkit_web_view_new                  func() uintptr
	webkit_web_view_get_type             func() uintptr
	webkit_web_view_load_uri             func(uintptr, *byte)
	webkit_website_data_manager_get_type func() uintptr
	webkit_web_view_try_close            func(uintptr)
	webkit_web_view_load_html            func(uintptr, content uintptr, base_uri uintptr)
	webkit_web_view_evaluate_javascript  func(
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
	webkit_security_origin_new         func(protocol, host *byte, port uint16) uintptr
	webkit_security_origin_new_for_uri func(uintptr) uintptr

	//Navigation
	webkit_navigation_policy_decision_get_navigation_action func(uintptr) uintptr
	webkit_navigation_action_get_request                    func(uintptr) uintptr
	webkit_uri_request_get_uri                              func(uintptr) uintptr

	//Network
	webkit_network_session_new                         func(dataDir, cacheDir *byte) uintptr
	webkit_network_session_get_default                 func() uintptr
	webkit_network_session_is_ephemeral                func(uintptr) bool
	webkit_network_proxy_settings_new                  func(default_proxy_uri, ignore_hosts unsafe.Pointer) uintptr
	webkit_network_proxy_settings_free                 func(proxy_settings uintptr)
	webkit_network_session_set_proxy_settings          func(session, proxy_mode, proxy_settings uintptr)
	webkit_network_proxy_settings_add_proxy_for_scheme func(proxy_settings uintptr, scheme, proxy_uri unsafe.Pointer)
	webkit_authentication_request_is_for_proxy         func(request uintptr) int32
	webkit_network_proxy_settings_copy                 func(proxy_settings uintptr) uintptr

	//Context
	webkit_web_context_new                                 func() uintptr
	webkit_web_context_get_type                            func() uintptr
	webkit_web_context_initialize_notification_permissions func(
		context uintptr,
		allowedOrigins uintptr,
		disallowedOrigins uintptr,
	)

	//Settings
	webkit_web_view_get_settings                        func(webview uintptr) uintptr
	webkit_settings_set_user_agent                      func(settings uintptr, userAgent uintptr)
	webkit_settings_set_enable_web_audio                func(uintptr, bool)
	webkit_settings_set_enable_javascript               func(uintptr, bool)
	webkit_settings_set_javascript_can_access_clipboard func(uintptr, bool)
	webkit_settings_set_enable_webgl                    func(uintptr, bool)
	webkit_settings_set_enable_page_cache               func(uintptr, bool)
	webkit_settings_set_enable_smooth_scrolling         func(uintptr, bool)
	webkit_settings_set_auto_load_images                func(uintptr, bool)
	webkit_settings_set_hardware_acceleration_policy    func(uintptr, int)
	webkit_settings_set_enable_developer_extras         func(uintptr, bool)

	webkit_settings_set_allow_file_access_from_file_urls        func(uintptr, bool)
	webkit_settings_set_allow_universal_access_from_file_urls   func(uintptr, bool)
	webkit_settings_set_enable_2d_canvas_acceleration           func(uintptr, bool)
	webkit_settings_set_enable_hyperlink_auditing               func(uintptr, bool)
	webkit_settings_set_enable_media_stream                     func(uintptr, bool)
	webkit_settings_set_enable_media_capabilities               func(uintptr, bool)
	webkit_settings_set_enable_encrypted_media                  func(uintptr, bool)
	webkit_settings_set_enable_site_specific_quirks             func(uintptr, bool)
	webkit_settings_set_enable_resizable_text_areas             func(uintptr, bool)
	webkit_settings_set_enable_tabs_to_links                    func(uintptr, bool)
	webkit_settings_set_enable_back_forward_navigation_gestures func(uintptr, bool)
	webkit_settings_set_enable_write_console_messages_to_stdout func(uintptr, bool)

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

	purego.RegisterLibFunc(&webkit_web_view_new, lib, "webkit_web_view_new")
	purego.RegisterLibFunc(&webkit_web_view_get_type, lib, "webkit_web_view_get_type")
	purego.RegisterLibFunc(&webkit_web_view_load_uri, lib, "webkit_web_view_load_uri")
	purego.RegisterLibFunc(&webkit_website_data_manager_get_type, lib, "webkit_website_data_manager_get_type")
	purego.RegisterLibFunc(&webkit_web_view_try_close, lib, "webkit_web_view_try_close")
	purego.RegisterLibFunc(&webkit_web_view_load_html, lib, "webkit_web_view_load_html")
	purego.RegisterLibFunc(&webkit_web_view_evaluate_javascript, lib, "webkit_web_view_evaluate_javascript")

	//Security
	purego.RegisterLibFunc(&webkit_security_origin_new, lib, "webkit_security_origin_new")
	purego.RegisterLibFunc(&webkit_security_origin_new_for_uri, lib, "webkit_security_origin_new_for_uri")

	//Navigation
	purego.RegisterLibFunc(&webkit_navigation_policy_decision_get_navigation_action, lib, "webkit_navigation_policy_decision_get_navigation_action")
	purego.RegisterLibFunc(&webkit_navigation_action_get_request, lib, "webkit_navigation_action_get_request")
	purego.RegisterLibFunc(&webkit_uri_request_get_uri, lib, "webkit_uri_request_get_uri")

	//Context
	purego.RegisterLibFunc(&webkit_web_context_new, lib, "webkit_web_context_new")
	purego.RegisterLibFunc(&webkit_web_context_get_type, lib, "webkit_web_context_get_type")

	//Network Session
	purego.RegisterLibFunc(&webkit_network_session_new, lib, "webkit_network_session_new")
	purego.RegisterLibFunc(&webkit_network_session_get_default, lib, "webkit_network_session_get_default")
	purego.RegisterLibFunc(&webkit_network_session_is_ephemeral, lib, "webkit_network_session_is_ephemeral")
	purego.RegisterLibFunc(&webkit_web_context_initialize_notification_permissions, lib, "webkit_web_context_initialize_notification_permissions")
	purego.RegisterLibFunc(&webkit_network_proxy_settings_new, lib, "webkit_network_proxy_settings_new")
	purego.RegisterLibFunc(&webkit_network_proxy_settings_free, lib, "webkit_network_proxy_settings_free")
	purego.RegisterLibFunc(&webkit_network_session_set_proxy_settings, lib, "webkit_network_session_set_proxy_settings")
	purego.RegisterLibFunc(&webkit_network_proxy_settings_add_proxy_for_scheme, lib, "webkit_network_proxy_settings_add_proxy_for_scheme")
	purego.RegisterLibFunc(&webkit_authentication_request_is_for_proxy, lib, "webkit_authentication_request_is_for_proxy")
	purego.RegisterLibFunc(&webkit_network_proxy_settings_copy, lib, "webkit_network_proxy_settings_copy")

	//Settings
	purego.RegisterLibFunc(&webkit_web_view_get_settings, lib, "webkit_web_view_get_settings")
	purego.RegisterLibFunc(&webkit_settings_set_user_agent, lib, "webkit_settings_set_user_agent")
	purego.RegisterLibFunc(&webkit_settings_set_enable_web_audio, lib, "webkit_settings_set_enable_webaudio")
	purego.RegisterLibFunc(&webkit_settings_set_enable_javascript, lib, "webkit_settings_set_enable_javascript")
	purego.RegisterLibFunc(&webkit_settings_set_javascript_can_access_clipboard, lib, "webkit_settings_set_javascript_can_access_clipboard")
	purego.RegisterLibFunc(&webkit_settings_set_enable_webgl, lib, "webkit_settings_set_enable_webgl")
	purego.RegisterLibFunc(&webkit_settings_set_enable_page_cache, lib, "webkit_settings_set_enable_page_cache")
	purego.RegisterLibFunc(&webkit_settings_set_enable_smooth_scrolling, lib, "webkit_settings_set_enable_smooth_scrolling")
	purego.RegisterLibFunc(&webkit_settings_set_auto_load_images, lib, "webkit_settings_set_auto_load_images")
	purego.RegisterLibFunc(&webkit_settings_set_hardware_acceleration_policy, lib, "webkit_settings_set_hardware_acceleration_policy")
	purego.RegisterLibFunc(&webkit_settings_set_enable_developer_extras, lib, "webkit_settings_set_enable_developer_extras")

	purego.RegisterLibFunc(&webkit_settings_set_allow_file_access_from_file_urls, lib, "webkit_settings_set_allow_file_access_from_file_urls")
	purego.RegisterLibFunc(&webkit_settings_set_allow_universal_access_from_file_urls, lib, "webkit_settings_set_allow_universal_access_from_file_urls")
	purego.RegisterLibFunc(&webkit_settings_set_enable_2d_canvas_acceleration, lib, "webkit_settings_set_enable_2d_canvas_acceleration")
	purego.RegisterLibFunc(&webkit_settings_set_enable_hyperlink_auditing, lib, "webkit_settings_set_enable_hyperlink_auditing")
	purego.RegisterLibFunc(&webkit_settings_set_enable_media_stream, lib, "webkit_settings_set_enable_media_stream")
	purego.RegisterLibFunc(&webkit_settings_set_enable_media_capabilities, lib, "webkit_settings_set_enable_media_capabilities")
	purego.RegisterLibFunc(&webkit_settings_set_enable_encrypted_media, lib, "webkit_settings_set_enable_encrypted_media")
	purego.RegisterLibFunc(&webkit_settings_set_enable_site_specific_quirks, lib, "webkit_settings_set_enable_site_specific_quirks")
	purego.RegisterLibFunc(&webkit_settings_set_enable_resizable_text_areas, lib, "webkit_settings_set_enable_resizable_text_areas")
	purego.RegisterLibFunc(&webkit_settings_set_enable_tabs_to_links, lib, "webkit_settings_set_enable_tabs_to_links")
	purego.RegisterLibFunc(&webkit_settings_set_enable_back_forward_navigation_gestures, lib, "webkit_settings_set_enable_back_forward_navigation_gestures")
	purego.RegisterLibFunc(&webkit_settings_set_enable_write_console_messages_to_stdout, lib, "webkit_settings_set_enable_write_console_messages_to_stdout")

	//permission
	purego.RegisterLibFunc(&webkit_permission_request_allow, lib, "webkit_permission_request_allow")
	purego.RegisterLibFunc(&webkit_notification_permission_request_get_type, lib, "webkit_notification_permission_request_get_type")
	purego.RegisterLibFunc(&webkit_geolocation_permission_request_get_type, lib, "webkit_geolocation_permission_request_get_type")

	//Content manager
	purego.RegisterLibFunc(&webkit_user_content_manager_new, lib, "webkit_user_content_manager_new")
	purego.RegisterLibFunc(&webkit_user_content_manager_register_script_message_handler, lib, "webkit_user_content_manager_register_script_message_handler")
	purego.RegisterLibFunc(&webkit_user_content_manager_register_script_message_handler_with_reply, lib, "webkit_user_content_manager_register_script_message_handler_with_reply")

}
