package convert

// ToString 将任意值安全转换为字符串。
func ToString(value any) string {
	text, ok := value.(string)
	if !ok {
		return ""
	}
	return text
}
