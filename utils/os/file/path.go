package osfile

import (
	"path/filepath"
	"strings"
)

// Path 返回规范化后的文件路径。
func Path(path string) string {
	return filepath.Clean(strings.TrimSpace(path))
}

// Resolve 返回文件绝对/相对路径的统一结果（相对路径会拼接 base）。
func Resolve(path string, base string) string {
	cleanPath := Path(path)
	if cleanPath == "" {
		return ""
	}
	if filepath.IsAbs(cleanPath) {
		return cleanPath
	}

	cleanBase := Path(base)
	if cleanBase == "" {
		return cleanPath
	}
	return filepath.Clean(filepath.Join(cleanBase, cleanPath))
}
