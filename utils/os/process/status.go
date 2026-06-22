package process

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"
)

// Alive 判断给定 PID 的进程是否仍存活。
func Alive(pid int) bool {
	if pid <= 0 {
		return false
	}

	target, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	err = target.Signal(syscall.Signal(0))
	if err == nil {
		return !zombie(pid)
	}
	// ESRCH 表示进程不存在，其他错误（如 EPERM）视为进程存在。
	var errno syscall.Errno
	if errors.As(err, &errno) && errno == syscall.ESRCH {
		return false
	}
	return !errors.Is(err, os.ErrProcessDone)
}

func zombie(pid int) bool {
	raw, err := os.ReadFile(fmt.Sprintf("/proc/%d/stat", pid))
	if err != nil {
		return false
	}

	stat := string(raw)
	end := strings.LastIndex(stat, ")")
	if end < 0 || len(stat) <= end+2 {
		return false
	}
	return stat[end+2] == 'Z'
}
