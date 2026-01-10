package webkit

import (
	"runtime"
	"unsafe"
)

func WebViewGetSettings(webView uintptr) uintptr {
	return webkit_web_view_get_settings(webView)
}

func SettingsSetUserAgent(settings uintptr, userAgent string) {
	userAgentC := cstring(userAgent)
	webkit_settings_set_user_agent(settings, userAgentC)
}

func SettingsSetEnableWebAudio(settings uintptr, enable bool) {
	webkit_settings_set_enable_web_audio(settings, enable)
}

func SettingsSetEnableJavascript(settings uintptr, enable bool) {
	webkit_settings_set_enable_javascript(settings, enable)
}

func SettingsSetJavascriptCanAccessClipboard(settings uintptr, enable bool) {
	webkit_settings_set_javascript_can_access_clipboard(settings, enable)
}

func SettingsSetEnableWebGL(settings uintptr, enable bool) {
	webkit_settings_set_enable_webgl(settings, enable)
}

func SettingsSetEnablePageCache(settings uintptr, enable bool) {
	webkit_settings_set_enable_page_cache(settings, enable)
}

func SettingsSetEnableSmoothScrolling(settings uintptr, enable bool) {
	webkit_settings_set_enable_smooth_scrolling(settings, enable)
}

func SettingsSetAutoLoadImages(settings uintptr, enable bool) {
	webkit_settings_set_auto_load_images(settings, enable)
}

func SettingsSetEnableDeveloperExtras(settings uintptr, enable bool) {
	webkit_settings_set_enable_developer_extras(settings, enable)
}

/*
## List of types
WEBKIT_HARDWARE_ACCELERATION_POLICY_ALWAYS = 0
WEBKIT_HARDWARE_ACCELERATION_POLICY_NEVER = 1
*/
func SettingsSetHardwareAccelerationPolicy(settings uintptr, policy int) {
	webkit_settings_set_hardware_acceleration_policy(settings, policy)
}

func SettingsSetAllowFileAccessFromFileUrls(settings uintptr, allow bool) {
	webkit_settings_set_allow_file_access_from_file_urls(settings, allow)
}

func SettingsSetAllowUniversalAccessFromFileUrls(settings uintptr, allow bool) {
	webkit_settings_set_allow_universal_access_from_file_urls(settings, allow)
}

func SettingsSetEnable2dCanvasAcceleration(settings uintptr, enable bool) {
	webkit_settings_set_enable_2d_canvas_acceleration(settings, enable)
}

func SettingsSetEnableHyperlinkAuditing(settings uintptr, enable bool) {
	webkit_settings_set_enable_hyperlink_auditing(settings, enable)
}

func SettingsSetEnableMediaStream(settings uintptr, enable bool) {
	webkit_settings_set_enable_media_stream(settings, enable)
}

func SettingsSetEnableMediaCapabilities(settings uintptr, enable bool) {
	webkit_settings_set_enable_media_capabilities(settings, enable)
}

func SettingsSetEnableEncryptedMedia(settings uintptr, enable bool) {
	webkit_settings_set_enable_encrypted_media(settings, enable)
}

func SettingsSetEnableSiteSpecificQuirks(settings uintptr, enable bool) {
	webkit_settings_set_enable_site_specific_quirks(settings, enable)
}

func SettingsSetEnableResizableTextAreas(settings uintptr, enable bool) {
	webkit_settings_set_enable_resizable_text_areas(settings, enable)
}

func SettingsSetEnableTabsToLinks(settings uintptr, enable bool) {
	webkit_settings_set_enable_tabs_to_links(settings, enable)
}

func SettingsSetEnableBackForwardNavigationGestures(settings uintptr, enable bool) {
	webkit_settings_set_enable_back_forward_navigation_gestures(settings, enable)
}

func SettingsSetEnableWriteConsoleMessagesToStdout(settings uintptr, enable bool) {
	webkit_settings_set_enable_write_console_messages_to_stdout(settings, enable)
}

/*
===========================
Proxy
===========================
*/

/*
Create a new WebKitNetworkProxySettings with the given default_proxy_uri and ignore_hosts.

The default proxy URI will be used for any URI that doesn’t match ignore_hosts, and doesn’t match any of the schemes added with webkit_network_proxy_settings_add_proxy_for_scheme(). If default_proxy_uri starts with “socks://”, it will be treated as referring to all three of the socks5, socks4a, and socks4 proxy types.

ignore_hosts is a list of hostnames and IP addresses that the resolver should allow direct connections to. Entries can be in one of 4 formats: A hostname, such as “example.com”, “.example.com”, or “*.example.com”, any of which match “example.com” or any subdomain of it. An IPv4 or IPv6 address, such as “192.168.1.1”, which matches only that address. A hostname or IP address followed by a port, such as “example.com:80”, which matches whatever the hostname or IP address would match, but only for URLs with the (explicitly) indicated port. In the case of an IPv6 address, the address part must appear in brackets: “[::1]:443” An IP address range, given by a base address and prefix length, such as “fe80::/10”, which matches any address in that range.

Note that when dealing with Unicode hostnames, the matching is done against the ASCII form of the name. Also note that hostname exclusions apply only to connections made to hosts identified by name, and IP address exclusions apply only to connections made to hosts identified by address. That is, if example.com has an address of 192.168.1.1, and ignore_hosts contains only “192.168.1.1”, then a connection to “example.com” will use the proxy, and a connection to 192.168.1.1” will not.

Available since: 2.16
*/
func NetworkProxySettingsNew(defaultProxy string, ignoreHosts []string) uintptr {
	cDefaultProxy := append([]byte(defaultProxy), 0)

	cHostsBytes := make([][]byte, len(ignoreHosts))
	cIgnoreHosts := make([]unsafe.Pointer, len(ignoreHosts)+1)

	for i, host := range ignoreHosts {
		cHostsBytes[i] = append([]byte(host), 0)
		cIgnoreHosts[i] = unsafe.Pointer(&cHostsBytes[i][0])
	}
	cIgnoreHosts[len(ignoreHosts)] = nil

	runtime.KeepAlive(cDefaultProxy)
	runtime.KeepAlive(cHostsBytes)
	runtime.KeepAlive(cIgnoreHosts)

	res := webkit_network_proxy_settings_new(
		unsafe.Pointer(&cDefaultProxy[0]),
		unsafe.Pointer(&cIgnoreHosts[0]),
	)

	return res
}

/*
Free the WebKitNetworkProxySettings.

Available since: 2.16
*/
func NetworkProxySettingsFree(settings uintptr) {
	webkit_network_proxy_settings_free(settings)
}

/*
Set the network proxy settings to be used by connections started in session session.

By default WEBKIT_NETWORK_PROXY_MODE_DEFAULT is used, which means that the system settings will be used (g_proxy_resolver_get_default()). If you want to override the system default settings, you can either use WEBKIT_NETWORK_PROXY_MODE_NO_PROXY to make sure no proxies are used at all, or WEBKIT_NETWORK_PROXY_MODE_CUSTOM to provide your own proxy settings. When proxy_mode is WEBKIT_NETWORK_PROXY_MODE_CUSTOM proxy_settings must be a valid WebKitNetworkProxySettings; otherwise, proxy_settings must be NULL.

Available since: 2.40
*/
func NetworkSessionSetProxySettings(session uintptr, proxy_mode NetworkProxyMode, proxy_settings uintptr) {
	webkit_network_session_set_proxy_settings(session, uintptr(proxy_mode), proxy_settings)
}

/*
Adds a URI-scheme-specific proxy.

URIs whose scheme matches uri_scheme will be proxied via proxy_uri. As with the default proxy URI, if proxy_uri starts with “socks://”, it will be treated as referring to all three of the socks5, socks4a, and socks4 proxy types.

Available since: 2.16
*/
func NetworkProxySettingAddProxyForScheme(proxy_settings uintptr, scheme ProxyScheme, proxy_uri string) {
	cScheme := append([]byte(scheme), 0)
	cProxyUri := append([]byte(proxy_uri), 0)

	webkit_network_proxy_settings_add_proxy_for_scheme(proxy_settings, unsafe.Pointer(&cScheme[0]), unsafe.Pointer(&cProxyUri[0]))

	runtime.KeepAlive(cScheme)
	runtime.KeepAlive(cProxyUri)
}

/*
Determine whether the authentication challenge is associated with a proxy server.

Determine whether the authentication challenge is associated with a proxy server rather than an “origin” server.

Available since: 2.2
*/
func AuthenticationRequestIsForProxy(request uintptr) bool {
	ret := webkit_authentication_request_is_for_proxy(request)

	if ret == 1 {
		return true
	}

	return false
}

/*
Make a copy of the WebKitNetworkProxySettings.

Available since: 2.16
*/
func NetworkProxySettingsCopy(proxy_settings uintptr) uintptr {
	return webkit_network_proxy_settings_copy(proxy_settings)
}
