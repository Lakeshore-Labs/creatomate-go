package creatomate

import (
	"encoding/json"
	"fmt"
)

// DebugJSON prints a nicely formatted JSON representation of any value
func DebugJSON(name string, value interface{}) {
	jsonBytes, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		fmt.Printf("%s: Error marshaling JSON: %v\n", name, err)
		return
	}
	fmt.Printf("%s:\n%s\n", name, jsonBytes)
}