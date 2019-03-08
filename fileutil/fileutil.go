// Copyright 2019 go-utils Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package fileutil

import (
	"strings"
	"os"
	"io"
)

// 检查路径前缀是否包含"/"，如果不包含PrefixPath返回"/"+path
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

// 如果文件或目录存在，Exist 返回true
func Exist(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}

// 将src文件内容拷贝到dest文件
func CopyFile(dest string, src string) error {
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
