package util

import (
	"os"
	"path/filepath"
)

/**
 * 检查path所在目录是否存在，如果不存在自动创建
 */
func MkdirPath(input string) (err error) {
	p := filepath.Dir(input)
	if _, tmp := os.Stat(p); tmp != nil && !os.IsExist(tmp) {
		err = os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return
		}
	}
	return
}