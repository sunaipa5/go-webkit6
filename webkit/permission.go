package webkit

func PermissionRequestAllow(request uintptr) {
	webkit_permission_request_allow(request)
}

func NotificationPermissionRequestGetType() uintptr {
	return webkit_notification_permission_request_get_type()
}

func GeolocationPermissionRequestGetType() uintptr {
	return webkit_geolocation_permission_request_get_type()
}
