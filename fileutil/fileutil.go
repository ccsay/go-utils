package fileutil

import (
	"strings"
	"os"
	"io"
)

// 处理路径前缀
func PrefixPath(path string) string {
	if strings.HasPrefix(path, "/") {
		return path
	}
	return "/" + path
}

// 处理路径后缀
func SuffixPath(path string) string {
	if strings.HasSuffix(path, "/") {
		return path
	}
	return path + "/"
}

// 判断文件夹或文件是否存在
func PathOrFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 将src文件内容拷贝到dest文件
func copyFile(dest string, src string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer srcFile.Close()
	defer destFile.Close()
	_, err = io.Copy(destFile, srcFile)
	return err
}
