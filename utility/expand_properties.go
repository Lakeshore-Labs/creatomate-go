package utility

import "reflect"

// ExpandProperties expands nested properties and keyframes
func ExpandProperties(properties map[string]interface{}) map[string]interface{} {
	expanded := make(map[string]interface{})
	
	for key, value := range properties {
		// For now, just copy the properties as-is
		// In a full implementation, this would handle keyframes and nested properties
		expanded[key] = value
	}
	
	return expanded
}

// IsKeyframeArray checks if the value is an array of keyframes
func IsKeyframeArray(value interface{}) bool {
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Slice {
		return false
	}
	
	// Check if it's a slice of maps with "time" and "value" fields
	if rv.Len() > 0 {
		elem := rv.Index(0)
		if elem.Kind() == reflect.Map {
			// Simple check for keyframe structure
			if m, ok := elem.Interface().(map[string]interface{}); ok {
				_, hasTime := m["time"]
				_, hasValue := m["value"]
				return hasTime && hasValue
			}
		}
	}
	
	return false
}