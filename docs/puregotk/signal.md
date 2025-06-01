# Signal Connect

## Permission Request

Handle permission request.Notification, microphone etc.

```go

	permissionFunc := func(webview, request, userData uintptr) {
		fmt.Println("Permission request")
		fmt.Printf("Permission request pointer: 0x%x\n", request)

   // Compare request type
		if gobject.TypeCheckInstanceIsA((*gobject.TypeInstance)(unsafe.Pointer(request)), types.GType(webkit.NotificationPermissionRequestGetType())) {
			//Allow request
			webkit.PermissionRequestAllow(request)
		}
	}

	gobject.SignalConnect(webview, "permission-request", glib.NewCallback(&permissionFunc))
```

## Decide Policy

```go
	decideFunc := func(webview, decision uintptr, decisionType int32) {
		switch decisionType {
		case 0:
			fmt.Println("WEBKIT_POLICY_DECISION_TYPE_RESPONSE")
		case 1:
			fmt.Println("WEBKIT_POLICY_DECISION_TYPE_NAVIGATION_ACTION")
		case 2:
			fmt.Println("WEBKIT_POLICY_DECISION_TYPE_NEW_WINDOW_ACTION")
		case 3:
			fmt.Println("WEBKIT_POLICY_DECISION_TYPE_DOWNLOAD")
		default:
			fmt.Println("UNKNOWN TYPE:", decisionType)
		}
	}

	gobject.SignalConnect(webview, "decide-policy", glib.NewCallback(&decideFunc))

```

## Inject js on dom ready

```go
	loadChangedFunc := func(webview, loadEvent, userData uintptr) {
		const WEBKIT_LOAD_FINISHED = 3
		if loadEvent == WEBKIT_LOAD_FINISHED {
			fmt.Println("Load finished, injecting JS")

			webkit.WebViewEvaluateJavascript(webview, `alert("hello x1289");`, 0, 0, 0, 0, 0, 0)
		}
	}

	gobject.SignalConnect(webview, "load-changed", glib.NewCallback(&loadChangedFunc))
```

### Notification request

```go
notificationFunc := func(webview, request, userData uintptr) {
		fmt.Println("Notification request")
	}

	gobject.SignalConnect(webview, "show-notification", glib.NewCallback(&notificationFunc))
```
