package safe

// String : safely access a string pointer
func String(x interface{}) string {
	v, ok := x.(string)
	if ok {
		return v
	}
	return ""
}

// StringSlice : safely access a slice of strings
func StringSlice(x interface{}) []string {
	is, ok := x.([]interface{})
	if ok {
		ss := make([]string, len(is))
		for i, iv := range is {
			ss[i] = iv.(string)
		}
		return ss
	}
	return []string{}
}

// MapStringInferface : safely access a map string interface
func MapStringInferface(x interface{}) map[string]interface{} {
	msi := make(map[string]interface{})
	mii, ok := x.(map[interface{}]interface{})
	if ok {
		for k, v := range mii {
			msi[k.(string)] = v
		}
	}

	return msi
}
