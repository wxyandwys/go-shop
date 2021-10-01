package config

func CodeJSON(code int, message string, data map[string]interface{}) (map[string]interface{}) {
	m := map[string]interface{}{}
	m["code"] = code
	m["message"] = message
	m["data"] = data
	return m
}