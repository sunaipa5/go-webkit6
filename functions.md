# WebkitGTK-6.0

## General

- webkit_web_view_new
- webkit_web_view_get_type
- webkit_web_view_load_uri
- webkit_website_data_manager_get_type
- webkit_web_view_try_close
- webkit_web_view_load_html
- webkit_web_view_evaluate_javascript

## Security

- webkit_security_origin_new
- webkit_security_origin_new_for_uri

## Navigation

- webkit_navigation_policy_decision_get_navigation_action
- webkit_navigation_action_get_request
- webkit_uri_request_get_uri

## Context

- webkit_web_context_new
- webkit_web_context_get_type

## Network Session

- webkit_network_session_new
- webkit_network_session_get_default
- webkit_network_session_is_ephemeral
- webkit_network_proxy_settings_new
- webkit_network_proxy_settings_free
- webkit_network_session_set_proxy_settings
- webkit_network_proxy_settings_add_proxy_for_scheme
- webkit_authentication_request_is_for_proxy
- webkit_network_proxy_settings_copy

## Settings

- webkit_web_view_get_settings
- webkit_settings_set_user_agent
- webkit_settings_set_enable_webaudio
- webkit_settings_set_enable_javascript
- webkit_settings_set_javascript_can_access_clipboard
- webkit_settings_set_enable_webgl
- webkit_settings_set_enable_page_cache
- webkit_settings_set_enable_smooth_scrolling
- webkit_settings_set_auto_load_images
- webkit_settings_set_hardware_acceleration_policy
- webkit_settings_set_enable_developer_extras
- webkit_settings_set_allow_file_access_from_file_urls
- webkit_settings_set_allow_universal_access_from_file_urls
- webkit_settings_set_enable_2d_canvas_acceleration
- webkit_settings_set_enable_hyperlink_auditing
- webkit_settings_set_enable_media_stream
- webkit_settings_set_enable_media_capabilities
- webkit_settings_set_enable_encrypted_media
- webkit_settings_set_enable_site_specific_quirks
- webkit_settings_set_enable_resizable_text_areas
- webkit_settings_set_enable_tabs_to_links
- webkit_settings_set_enable_back_forward_navigation_gestures
- webkit_settings_set_enable_write_console_messages_to_stdout

## Permission

- webkit_permission_request_allow
- webkit_notification_permission_request_get_type
- webkit_geolocation_permission_request_get_type
- webkit_web_context_initialize_notification_permissions

# JavaScriptCoreGTK-6.0

## Value

- jsc_value_is_array
- jsc_value_is_array_buffer
- jsc_value_is_boolean
- jsc_value_is_constructor
- jsc_value_is_function
- jsc_value_is_null
- jsc_value_is_number
- jsc_value_is_object
- jsc_value_is_string
- jsc_value_is_typed_array
- jsc_value_is_undefined

### Converters

- jsc_value_to_boolean
- jsc_value_to_double
- jsc_value_to_int32
- jsc_value_to_json
- jsc_value_to_string
- jsc_value_to_string_as_bytes
