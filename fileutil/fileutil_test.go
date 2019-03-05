package fileutil

import (
	"testing"
	"time"
	"fmt"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathOrFileExist(tt.args.path); got != tt.want {
				t.Errorf("PathOrFileExist() = %v, want %v", got, tt.want)
			}
		})
	}
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
