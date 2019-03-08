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
	"testing"
	"time"
	"fmt"
	"strings"
	"runtime"
	"path"
)

func TestPrefixPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrefixPath(tt.args.path); got != tt.want {
				t.Errorf("PrefixPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuffixPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SuffixPath(tt.args.path); got != tt.want {
				t.Errorf("SuffixPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPathOrFileExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "dir",
			args: args{
				path: getProjectRootPath(),
			},
			want: true,
		}, {
			name: "file",
			args: args{
				path: getProjectRootPath() + "/common/utils/utils.go",
			},
			want: true,
		}, {
			name: "file is nil",
			args: args{
				path: getProjectRootPath() + "/common/utils/utils1.go",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.path)
			if got := PathOrFileExist(tt.args.path); got != tt.want {
				t.Errorf("PathOrFileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getProjectRootPath() string {
	currentPath := getCurrentPath()
	return strings.Replace(currentPath, "/common/utils", "", 1)
}

func getCurrentPath() string {
	// skip：0.表示调用者本身，获取的是当前文件名
	// skip：1.表示调用者的调用者，获取的是源头调用文件名
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}


func Test_copyFile(t *testing.T) {
	type args struct {
		dest string
		src  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				src: "/Users/lcl/file/教程/go/老男孩go/L004-Go语言/01 Go开发1期 day4 课后作业讲解01.mp4.avi",
				dest:  "./1.mp4.avi",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		start := time.Now()
		t.Run(tt.name, func(t *testing.T) {
			if err := copyFile(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("copyFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		fmt.Println(time.Since(start))
	}
}

func Test_copyFile1(t *testing.T) {
	type args struct {
		dest string
		src  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				src: "/Users/lcl/file/教程/go/老男孩go/L004-Go语言/01 Go开发1期 day4 课后作业讲解01.mp4.avi",
				dest:  "./1.mp4.avi",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		start := time.Now()
		t.Run(tt.name, func(t *testing.T) {
			if err := copyFile1(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("copyFile1() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		fmt.Println(time.Since(start))
	}
}
