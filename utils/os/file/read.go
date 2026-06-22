package osfile

import (
	"os"
)

// Read 读取指定路径文件内容。
func Read(path string) ([]byte, error) {
	return os.ReadFile(Path(path))
}
