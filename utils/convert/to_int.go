package convert

// ToInt 将任意值安全转换为 int。
func ToInt(value any) int {
	switch typedValue := value.(type) {
	case int:
		return typedValue
	case int32:
		return int(typedValue)
	case int64:
		return int(typedValue)
	case float32:
		return int(typedValue)
	case float64:
		return int(typedValue)
	default:
		return 0
	}
}
