package javascriptcore

func ValueIsArray(jscValue uintptr) bool {
	return jscValueIsArray(jscValue) != 0
}

func ValueIsArrayBuffer(jscValue uintptr) bool {
	return jscValueIsArrayBuffer(jscValue) != 0
}

func ValueIsBoolean(jscValue uintptr) bool {
	return jscValueIsBoolean(jscValue) != 0
}

func ValueIsConstructor(jscValue uintptr) bool {
	return jscValueIsConstructor(jscValue) != 0
}

func ValueIsFunction(jscValue uintptr) bool {
	return jscValueIsFunction(jscValue) != 0
}

func ValueIsNull(jscValue uintptr) bool {
	return jscValueIsNull(jscValue) != 0
}

func ValueIsNumber(jscValue uintptr) bool {
	return jscValueIsNumber(jscValue) != 0
}

func ValueIsObject(jscValue uintptr) bool {
	return jscValueIsObject(jscValue) != 0
}

func ValueIsString(jscValue uintptr) bool {
	return jscValueIsString(jscValue) != 0
}

func ValueIsTypedArray(jscValue uintptr) bool {
	return jscValueIsTypedArray(jscValue) != 0
}

func ValueIsUndefined(jscValue uintptr) bool {
	return jscValueIsUndefined(jscValue) != 0
}

//============================
// Converters
//============================

func ValueToBoolean(jscValue uintptr) bool {
	return jscValueToBoolean(jscValue) != 0
}

func ValueToDouble(jscValue uintptr) float64 {
	return jscValueToDouble(jscValue)
}

func ValueToInt32(jscValue uintptr) int32 {
	return jscValueToInt32(jscValue)
}

func ValueToString(jscValue uintptr) string {
	ptr := jscValueToString(jscValue)
	if ptr == 0 {
		return ""
	}

	str := goString(ptr)

	return str
}

func ValueToJson(jscValue uintptr, indent uint32) string {
	ptr := jscValueToJson(jscValue, indent)
	if ptr == 0 {
		return ""
	}

	str := goString(ptr)

	return str
}
