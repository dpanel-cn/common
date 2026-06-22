package osfile

import (
	"errors"
	"os"
)

// Remove 删除指定文件；ignoreNotExist=true 时文件不存在不报错。
func Remove(path string, ignoreNotExist bool) error {
	err := os.Remove(Path(path))
	if err != nil && !(ignoreNotExist && errors.Is(err, os.ErrNotExist)) {
		return err
	}
	return nil
}
