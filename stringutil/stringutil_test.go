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
package stringutil

import (
	"fmt"
	"testing"
	"time"
)

func TestTrim(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				value: "  s  ",
			},
			want: "s",
		}, {
			name: "left",
			args: args{
				value: "  s",
			},
			want: "s",
		}, {
			name: "right",
			args: args{
				value: "s     ",
			},
			want: "s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.args.value); got != tt.want {
				t.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt64(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "ok",
			args: args{
				value: "123456",
			},
			want: int64(123456),
		}, {
			name: "err",
			args: args{
				value: "haha",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt64(tt.args.value); got != tt.want {
				t.Errorf("StringToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				length: 1000000,
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			got := RandomString(tt.args.length)
			fmt.Println(got)
			fmt.Println("ex", time.Now().Sub(start))
		})
	}
}
