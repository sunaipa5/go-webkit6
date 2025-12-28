package javascriptcore

import (
	"os"

	"github.com/jwijenbergh/purego"
)

var (
	//jsc_value_to_boolean
	jscValueToBoolean func(uintptr) int32
	//jsc_value_to_double
	jscValueToDouble func(uintptr) float64
	//jsc_value_to_int32
	jscValueToInt32 func(uintptr) int32
	//jsc_value_to_json
	jscValueToJson func(uintptr, uint32) uintptr
	//jsc_value_to_string
	jscValueToString func(uintptr) uintptr
	//jsc_value_to_string_as_bytes
	jscValueToStringAsBytes func(uintptr) uintptr

	jscValueIsArray       func(uintptr) int32
	jscValueIsArrayBuffer func(uintptr) int32
	jscValueIsBoolean     func(uintptr) int32
	jscValueIsConstructor func(uintptr) int32
	jscValueIsFunction    func(uintptr) int32
	jscValueIsNull        func(uintptr) int32
	jscValueIsNumber      func(uintptr) int32
	jscValueIsObject      func(uintptr) int32
	jscValueIsString      func(uintptr) int32
	jscValueIsTypedArray  func(uintptr) int32
	jscValueIsUndefined   func(uintptr) int32
)

// You can change the library location with ldflag or set the JAVASCRIPTCOREGTK_PATH environment variable.
// -X main.JavaScriptCoreGtkPath=/custom/path/libjavascriptcoregtk-6.0.so.1
// Don't change this variable while the application is running
var JavaScriptCoreGtkPath string

func init() {
	if JavaScriptCoreGtkPath == "" {
		if env := os.Getenv("JAVASCRIPTCOREGTK_PATH"); env != "" {
			JavaScriptCoreGtkPath = env
		} else {
			JavaScriptCoreGtkPath = "libjavascriptcoregtk-6.0.so.1"
		}
	}

	lib, err := purego.Dlopen(JavaScriptCoreGtkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	//============================
	// Value
	//============================

	purego.RegisterLibFunc(&jscValueIsArray, lib, "jsc_value_is_array")
	purego.RegisterLibFunc(&jscValueIsArrayBuffer, lib, "jsc_value_is_array_buffer")
	purego.RegisterLibFunc(&jscValueIsBoolean, lib, "jsc_value_is_boolean")
	purego.RegisterLibFunc(&jscValueIsConstructor, lib, "jsc_value_is_constructor")
	purego.RegisterLibFunc(&jscValueIsFunction, lib, "jsc_value_is_function")
	purego.RegisterLibFunc(&jscValueIsNull, lib, "jsc_value_is_null")
	purego.RegisterLibFunc(&jscValueIsNumber, lib, "jsc_value_is_number")
	purego.RegisterLibFunc(&jscValueIsObject, lib, "jsc_value_is_object")
	purego.RegisterLibFunc(&jscValueIsString, lib, "jsc_value_is_string")
	purego.RegisterLibFunc(&jscValueIsTypedArray, lib, "jsc_value_is_typed_array")
	purego.RegisterLibFunc(&jscValueIsUndefined, lib, "jsc_value_is_undefined")

	//Converters
	purego.RegisterLibFunc(&jscValueToBoolean, lib, "jsc_value_to_boolean")
	purego.RegisterLibFunc(&jscValueToDouble, lib, "jsc_value_to_double")
	purego.RegisterLibFunc(&jscValueToInt32, lib, "jsc_value_to_int32")
	purego.RegisterLibFunc(&jscValueToJson, lib, "jsc_value_to_json")
	purego.RegisterLibFunc(&jscValueToString, lib, "jsc_value_to_string")
	purego.RegisterLibFunc(&jscValueToStringAsBytes, lib, "jsc_value_to_string_as_bytes")

}
