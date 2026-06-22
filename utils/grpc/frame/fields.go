package frame

import (
	"fmt"
	"strings"

	convert "github.com/dpanel-cn/common/utils/convert"
)

// FirstString returns the first non-empty string from frame field candidates.
func FirstString(values ...any) string {
	for _, value := range values {
		text := strings.TrimSpace(convert.ToString(value))
		if text != "" {
			return text
		}
	}
	return ""
}

// BoolField reads a required boolean field from a gRPC Struct frame.
func BoolField(frame map[string]any, key string, label string) (bool, error) {
	value, ok := frame[key]
	if !ok {
		return false, fmt.Errorf("%s %s is required", label, key)
	}
	typed, ok := value.(bool)
	if !ok {
		return false, fmt.Errorf("%s %s must be boolean", label, key)
	}
	return typed, nil
}

// MapField reads an optional object field from a gRPC Struct frame.
func MapField(frame map[string]any, key string, label string) (map[string]any, error) {
	value, ok := frame[key]
	if !ok || value == nil {
		return map[string]any{}, nil
	}
	typed, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%s %s must be object", label, key)
	}
	return typed, nil
}
