package utility

import (
	"strings"
	"unicode"
)

// CamelToSnakeCase converts camelCase to snake_case
func CamelToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}

// SnakeToCamelCase converts snake_case to camelCase
func SnakeToCamelCase(s string) string {
	var result strings.Builder
	upperNext := false
	for _, r := range s {
		if r == '_' {
			upperNext = true
		} else {
			if upperNext {
				result.WriteRune(unicode.ToUpper(r))
				upperNext = false
			} else {
				result.WriteRune(r)
			}
		}
	}
	return result.String()
}

// TransformObjectKeysToSnake converts map keys from camelCase to snake_case
func TransformObjectKeysToSnake(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	
	switch v := obj.(type) {
	case map[string]interface{}:
		for key, value := range v {
			result[CamelToSnakeCase(key)] = value
		}
	default:
		// Use reflection for structs
		// This is a simplified version - in production would use reflection
		// For now, the caller should handle struct to map conversion
	}
	
	return result
}

// TransformObjectKeysToCamel converts map keys from snake_case to camelCase
func TransformObjectKeysToCamel(obj map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range obj {
		result[SnakeToCamelCase(key)] = value
	}
	return result
}