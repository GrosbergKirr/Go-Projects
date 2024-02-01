package tools

func TypeofObject(variable interface{}) string {
	switch variable.(type) {
	case int:
		return "int"
	case float32:
		return "float32"
	case bool:
		return "boolean"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
