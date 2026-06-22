package convert

import (
	"fmt"
	"reflect"
	"time"

	"google.golang.org/protobuf/types/known/structpb"
)

// ToStruct 将 map 安全转换为 protobuf Struct。
func ToStruct(payload map[string]any) *structpb.Struct {
	frame, err := structpb.NewStruct(normalizeMap(payload))
	if err == nil {
		return frame
	}

	// 理论上不会进入此兜底分支。
	fallback, _ := structpb.NewStruct(map[string]any{
		"ok":    false,
		"error": fmt.Sprintf("build struct frame failed: %v", err),
		"ts":    time.Now().Unix(),
	})
	return fallback
}

// normalizeMap 将 map 中值递归规范为 structpb 支持的数据类型。
func normalizeMap(payload map[string]any) map[string]any {
	if payload == nil {
		return map[string]any{}
	}
	normalized := make(map[string]any, len(payload))
	for key, value := range payload {
		normalized[key] = normalizeValue(value)
	}
	return normalized
}

// normalizeValue 将任意值递归规范为 structpb.NewValue 支持的类型集合。
func normalizeValue(value any) any {
	switch typed := value.(type) {
	case nil:
		return nil
	case bool, string, float64:
		return typed
	case float32:
		return float64(typed)
	case int:
		return float64(typed)
	case int8:
		return float64(typed)
	case int16:
		return float64(typed)
	case int32:
		return float64(typed)
	case int64:
		return float64(typed)
	case uint:
		return float64(typed)
	case uint8:
		return float64(typed)
	case uint16:
		return float64(typed)
	case uint32:
		return float64(typed)
	case uint64:
		return float64(typed)
	case []any:
		rows := make([]any, 0, len(typed))
		for _, item := range typed {
			rows = append(rows, normalizeValue(item))
		}
		return rows
	case map[string]any:
		return normalizeMap(typed)
	default:
		return normalizeByReflect(value)
	}
}

// normalizeByReflect 处理切片、数组、map、指针等动态类型。
func normalizeByReflect(value any) any {
	if value == nil {
		return nil
	}
	rv := reflect.ValueOf(value)
	for rv.IsValid() && (rv.Kind() == reflect.Interface || rv.Kind() == reflect.Pointer) {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		return nil
	}

	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		size := rv.Len()
		rows := make([]any, 0, size)
		for i := 0; i < size; i++ {
			rows = append(rows, normalizeValue(rv.Index(i).Interface()))
		}
		return rows
	case reflect.Map:
		rows := make(map[string]any, rv.Len())
		iter := rv.MapRange()
		for iter.Next() {
			key := fmt.Sprint(iter.Key().Interface())
			rows[key] = normalizeValue(iter.Value().Interface())
		}
		return rows
	default:
		return fmt.Sprint(rv.Interface())
	}
}
