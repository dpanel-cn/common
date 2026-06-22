package osfile

import (
	"os"
	"path/filepath"
)

// Write 写入指定文件内容；若父目录不存在会自动创建。
func Write(path string, content []byte, perm os.FileMode) error {
	cleanPath := Path(path)
	if perm == 0 {
		perm = 0o644
	}
	dirPerm := os.FileMode(0o755)
	if perm&0o077 == 0 {
		dirPerm = 0o700
	}

	// 写文件前先确保目录存在，避免启动时因目录缺失失败。
	if err := os.MkdirAll(filepath.Dir(cleanPath), dirPerm); err != nil {
		return err
	}
	return os.WriteFile(cleanPath, content, perm)
}
